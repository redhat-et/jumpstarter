---
title: Orin NX Devkit + DutLink
weight: 1
date: 2024-05-11
description: Manual for connecting the dutlink board to the Orin AGX devkit.
---

{{% pageinfo %}}
This is a graphical guide describing how to connect an NVIDIA Orin AGX Devkit
to the dutlink-board.
{{% /pageinfo %}}

This setup will use 4 USB connections to your host:

* Jumpstarter control
* Jumpstarter USB3 storage
* NVIDIA Flashing USBC or microUSB port
## Wiring table
<table class="table">
  <thead>
    <tr>
        <th scope="col">Name</th>
        <th scope="col">NX Devkit Connector</th>
        <th scope="col">Jumpstarter Connector</th>
        <th scope="col">Host connector</th>
        <th scope="col">Comments</th>
    </tr>
  </thead>
  <tbody>
    <tr>
        <th scope="row">GND</th>
        <td>(J14) pin 7</td>
        <td>(I/O) GND</td>
        <td></td>
        <td>Connecting signal ground</td>
    </tr>
    <tr>
        <th scope="row">/FORCE_REC</th>
        <td>(J14) pin 10</td>
        <td>(I/O) CTL_A</td>
        <td></td>
        <td>Force recovery mode signal (active low)</td>
    </tr>
    <tr>
        <th scope="row">SLEEP/WAKE</th>
        <td>(J14) pin 12</td>
        <td>(I/O) CTL_B</td>
        <td></td>
        <td>Power down [>10s], Power up [short] (active low)</td>
    </tr>
    <tr>
        <th scope="row">/SYS_RESET</th>
        <td>(J14) pin 8</td>
        <td>(I/O) RESET</td>
        <td></td>
        <td>Reset signal (active low)</td>
    </tr>
    <tr>
        <th scope="row">RCM</th>
        <td>(J5) USB-C or microUSB connector</td>
        <td></td>
        <td>USB</td>
        <td>NVIDIA Flashing interface for RCM</td>
    </tr>
    <tr>
        <th scope="row">DUT-STORAGE</th>
        <td>(J7-J6) USB 3.2 Gen1</td>
        <td>J8</td>
        <td></td>
        <td>USB storage attachment to DUT</td>
    </tr>
    <tr>
        <th scope="row">DUT-POWER</th>
        <td>POWER JACK</td>
        <td>J3</td>
        <td></td>
        <td>Power output for the DUT</td>
    </tr>
     <tr>
        <th scope="row">ETHERNET</th>
        <td>Ethernet</td>
        <td></td>
        <td></td>
        <td>Connect to a network where the host is also connected</td>
    </tr>
    <tr>
        <th scope="row">JUMPSTARTER</th>
        <td></td>
        <td>P1 USB-C</td>
        <td>USB</td>
        <td>Jumpstarter control USB bus, used by the jumpstarter software to talk to the
            dutlink-board</td>
    </tr>
    <tr>
        <th scope="row">HOST-STORAGE</th>
        <td></td>
        <td>J7 USB-B 3.0</td>
        <td>USB</td>
        <td>Host access to USB storage, used to write the USB disk</td>
    </tr>
    <tr>
        <th scope="row">DISK</th>
        <td></td>
        <td>J9 USB-A 3.0 </td>
        <td></td>
        <td>Connect a pen-drive or disk here. USB3.1 Gen1 (5Gbps recommended, Gen2 10Gbps don't work well yet)</td>
    </tr>
    <tr>
        <th scope="row">POWER-4-DUT</th>
        <td></td>
        <td>J2 BARREL JACK POWER</td>
        <td></td>
        <td>Connect the Orin NX Devkit power adapter here</td>
    </tr>
  </tbody>
</table>

## Troubleshooting
### My console doesn't show anything
See the [Console access](#console-access) section and clear any previous configuration for the USB console.


### The system won't boot from the USB disk
You need to go into the UEFI BIOS and change the boot order to setup "new devices"
as the first boot option. Make sure that the USB devices is found.

Make sure that you are not using a USB3.1 Gen2 device (10Gbps), as this is not supported yet.

## Power sequencing

The power sequencing settings are:

`jumpstarter set-config device-id power_on p1`

`jumpstarter set-config device-id power_off p0`

`jumpstarter set-config device-id power_rescue p1,aL,rL,w1,rZ,w1,aZ`

This means that we will use the power control directly, without using the power button as it's not necessary for this
board. The flashing mode is activated by setting force recovery signal to low `aL`, then asserting reset `rL`, and then
waiting for 1 second `w1`, then releasing the reset signal `rZ`, waiting another second `w1`, and then releasing the
force recovery signal `aZ`.

## Console access
The Orin NX Devkit the UEFI console only on the 40pin port, so any necessary UEFI settings must be performed
on that port first.

You must clear the previous usb console settings if this board was used with an AGX
{{< highlight "" >}}
$ jumpstarter set-usb-console orin-nx-00 ""
{{< /highlight >}}

