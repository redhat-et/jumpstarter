package jumpstarter_board

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

const BASE_DISKSBYID = "/dev/disk/by-id/"

const BLOCK_SIZE = 32 * 1024 * 1024

const WAIT_TIME_USB_STORAGE = 2 * time.Second

type StorageTarget int

const (
	HOST StorageTarget = iota
	DUT
	OFF
)

func (d *JumpstarterDevice) SetDiskImage(path string) error {

	fmt.Print("Detecting USB storage device and connecting to host: ")
	diskPath, err := d.detectStorageDevice()
	if err != nil {
		return fmt.Errorf("SetDiskImage: %w", err)
	}
	fmt.Println("done")

	fmt.Printf("%s -> %s: \n", path, diskPath)

	if err := writeImageToDisk(path, diskPath); err != nil {
		return fmt.Errorf("SetDiskImage: %w", err)
	}

	fmt.Print("Connecting storage to device under test.. ")
	if err := d.connectStorageTo(OFF); err != nil {
		return fmt.Errorf("SetDiskImage: %w", err)
	}

	return nil
}

func (d *JumpstarterDevice) AttachStorage(connected bool) error {
	var err error
	switch connected {
	case true:
		err = d.connectStorageTo(DUT)
	case false:
		err = d.connectStorageTo(OFF)
	}
	if err != nil {
		return fmt.Errorf("ConnectDiskImage(%v): %w", connected, err)
	}
	return nil
}

func (d *JumpstarterDevice) connectStorageTo(target StorageTarget) error {
	if err := d.ensureSerial(); err != nil {
		return fmt.Errorf("connectStorageTo: %w", err)
	}

	if err := d.exitConsole(); err != nil {
		return fmt.Errorf("connectStorageTo: %w", err)
	}

	var cmd string
	var response string
	switch target {
	case HOST:
		cmd = "host"
		response = "connected to host"
	case DUT:
		cmd = "dut"
		response = "connected to device"
	case OFF:
		cmd = "off"
		response = "storage disconnected"
	default:
		return fmt.Errorf("connectStorageTo: invalid target %v", target)
	}

	if err := d.sendAndExpect("storage "+cmd, response); err != nil {
		return fmt.Errorf("connectStorageTo(%q): %w", cmd, err)
	}
	return nil
}

func scanForStorageDevices(prefix string) (*mapset.Set[string], error) {

	diskSet := mapset.NewSet[string]()

	err := filepath.Walk(BASE_DISKSBYID, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "devices" {
			return nil
		}

		if info.Mode()&os.ModeSymlink != 0 {
			baseName := filepath.Base(path)
			if strings.HasPrefix(baseName, prefix) {
				diskSet.Add(path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("scanForStorageDevices: %w", err)
	}

	return &diskSet, nil
}

func (d *JumpstarterDevice) detectStorageDevice() (string, error) {
	if err := d.connectStorageTo(OFF); err != nil {
		return "", fmt.Errorf("detectStorageDevice: %w", err)
	}
	time.Sleep(WAIT_TIME_USB_STORAGE)

	diskSetOff, err := scanForStorageDevices("usb-")
	if err != nil {
		return "", fmt.Errorf("detectStorageDevice: %w", err)
	}

	if err := d.connectStorageTo(HOST); err != nil {
		return "", fmt.Errorf("detectStorageDevice: %w", err)
	}
	time.Sleep(WAIT_TIME_USB_STORAGE)

	diskSetOn, err := scanForStorageDevices("usb-")
	if err != nil {
		return "", fmt.Errorf("detectStorageDevice: %w", err)
	}

	newDiskSet := (*diskSetOn).Difference(*diskSetOff)
	if newDiskSet.Cardinality() != 1 {
		return "", fmt.Errorf("detectStorageDevice: expected one new disk, got %d, %v", newDiskSet.Cardinality(), newDiskSet)
	}

	diskPath, _ := newDiskSet.Pop()
	return diskPath, nil
}

func writeImageToDisk(imagePath string, diskPath string) error {
	inputFile, err := os.OpenFile(imagePath, os.O_RDONLY, 0666)
	if err != nil {
		return fmt.Errorf("writeImageToDisk: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile(diskPath, os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("writeImageToDisk: %w", err)
	}

	buffer := make([]byte, BLOCK_SIZE)
	bytesCopied := 0
	start := time.Now()
	for {
		n, err := inputFile.Read(buffer)
		if err != nil && err != io.EOF {
			outputFile.Close()
			return fmt.Errorf("writeImageToDisk: %w", err)
		}
		if n == 0 {
			break
		}
		if _, err := outputFile.Write(buffer[:n]); err != nil {
			outputFile.Close()
			return fmt.Errorf("writeImageToDisk: %w", err)
		}
		elapsed := time.Since(start)

		bytesCopied += n
		MBCopied := bytesCopied / 1024 / 1024
		MBPerSec := float64(MBCopied) / elapsed.Seconds()
		fmt.Printf("\r%d MB copied %.2f MB/s         ", MBCopied, MBPerSec)
	}
	outputFile.Close()
	fmt.Println()
	err = exec.Command("udisksctl", "power-off", "-b", diskPath).Run()
	if err != nil {
		return fmt.Errorf("writeImageToDisk: %w", err)
	}
	time.Sleep(WAIT_TIME_USB_STORAGE)
	return nil
}
