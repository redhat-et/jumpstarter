---
title: Orin AGX + jumpstarter
weight: 1
date: 2023-09-25
description: Manual for connecting the jumpstarter board to the Orin AGX devkit.
---

{{% pageinfo %}}
This is a graphical guide describing how to connect an NVIDIA Orin AGX Devkit
to the jumpstarter-board.
{{% /pageinfo %}}

This setup will use 4 USB connections to your host:

* Jumpstarter control
* Jumpstarter USB3 storage
* NVIDIA TOPO USB Controller
* NVIDIA Flashing USB port

The Orin AGX Devkit only exposes the UEFI and kernel serial console via the
micro USB port (also known as the NVIDIA TOPO USB controller).

Please see [The Orin AGX Devkit layout](https://developer.nvidia.com/embedded/learn/jetson-agx-orin-devkit-user-guide/developer_kit_layout.html) for more details.

To let jumpstarter know that it must look up for a specific usb serial port device
when trying to interact with the DUT console you will need to associate the
NVIDIA TOPO USB Console to your jumpstarter board using the
[usb-set-console](/docs/reference/#set-usb-console) command.

i.e. when the Orin AGX TOPO console shows up like this on the host:

{{< highlight "" >}}
[300810.229025] usb 1-1.3.1: new full-speed USB device number 27 using xhci_hcd
[300810.332797] usb 1-1.3.1: New USB device found, idVendor=0955, idProduct=7045, bcdDevice= 0.01
[300810.332799] usb 1-1.3.1: New USB device strings: Mfr=1, Product=2, SerialNumber=3
[300810.332800] usb 1-1.3.1: Product: Tegra On-Platform Operator
[300810.332801] usb 1-1.3.1: Manufacturer: NVIDIA
[300810.332801] usb 1-1.3.1: SerialNumber: TOPOD83B461B
{{< /highlight >}}

You should associate it to the jumpstarter board using the following command:

{{< highlight "" >}}
$ jumpstarter set-usb-console orin-agx-00 TOPOD83B461B-if01
{{< /highlight >}}


{{% imgproc "P9220014.jpg" Fit 1024x1024  %}}
Ethernet from the Orin AGX Devkit is connected to the a switch
where the host running jumpstarter has connectivity.

One USB is connected to the jumpstarter DUT storage connector (J8)

The Power USB-C connection is connected to J5 on the jumpstarter.
{{% /imgproc %}}

{{% imgproc "P9220041.jpg" Fit 1024x1024  %}}
In this picture we connected additional cables but we are only using
GND(black), /RESET(white), CTL_A(green) and CTL_B(blue)
{{% /imgproc %}}

{{% imgproc "P9220020.jpg" Fit 1024x1024 / %}}

{{% imgproc "P9220025.jpg" Fit 1024x1024 %}}
Pins 1,2,3,4 of the [Orin AGX devkit Automation Header J42](https://developer.nvidia.com/embedded/learn/jetson-agx-orin-devkit-user-guide/developer_kit_layout.html#automation-header-j42)

Must be connected to pins GND, CTL_A, CTL_B and /RESET of the jumpstarter board I/O connector.
{{% /imgproc %}}


{{% imgproc "P9220030.jpg" Fit 1024x1024 %}}

This USBC connection is used for flashing, and must be connected to the host via USB.

{{% /imgproc %}}

{{% imgproc "P9220032.jpg" Fit 1024x1024 / %}}

{{% imgproc "P9220037.jpg" Fit 1024x1024 %}}
Jumpstarter needs to be connected to the host via J7 and P1. And the Orin power adapter must be connected to J2.
{{% /imgproc %}}



{{% imgproc "P9220043.jpg" Fit 1024x1024 / %}}

