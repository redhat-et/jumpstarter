---
title: Connector reference
weight: 1
date: 2023-09-26
description: Connector reference for te jumpstarter-board

---

{{% imgproc "image.png" Fit 1024x1024 / %}}

## Device Under Test connections
### Power
<table class="table">
  <thead>
    <tr>
      <th scope="col">Connector</th>
      <th scope="col">Definition</th>
    </tr>
  </thead>
  <tbody>
   <tr>
      <th scope="row">J3</th>
      <td>J3 is a barrel jack connector that provides power output to the DUT. The source
            of the power comes from <b>J1</b>.<br/><br/>
        <b>Warning</b>: Do not use at the same time as J5 or J2.</td>
    </tr>
    <tr>
      <th scope="row">J5</th>
      <td>J5 is a USB-C connector that provides power output to the DUT. The source of the
power comes from <b>J2</b>, USB-PD negotiation is connected to the power adapter on J2. USB-PD is a standard that allows negotiation of the voltage and current. <br/><br/>
       <b>Warning:</b> Do not use at the same time as J3 or J1.</td>
    </tr>
  </tbody>
</table>

### Storage
<table class="table">
  <thead>
    <tr>
      <th scope="col">Connector</th>
      <th scope="col">Definition</th>
    </tr>
  </thead>
  <tbody>
   <tr>
      <th scope="row">J8</th>
      <td>J8 is a USB3 micro B connector for storage. This connector provides access to
          the storage device that is connected to <b>J9</b>.
            One possible cable you can use in this connector is this one: <a href="https://www.amazon.com/CableCreation-External-Compatible-Galaxy-Camera/dp/B074V3GD2S/">link</a></td>
    </tr>
  </tbody>
</table>



### I/O pins
The I/O pins connector provides a 3.3V digital interface to the DUT, or any surrounding
test hardware (like video signal generators, sensor emulators, or other elements).

Never use a HI/H/1 signal on an output when the target device is only 1.2, 1.8 or 2.5V, as this could damage the device.

Outputs are protected with a 100 ohm resistor which would avoid damage in most cases.

The V IO pin is provided to enable the use of voltage translation circuits if necessary.


<table class="table">
  <thead>
    <tr>
      <th scope="col">Pin</th>
      <th scope="col">Definition</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th scope="row">V IO</th>
      <td>Provides a voltage reference for a voltage translation circuit</td>
    </tr>
    <tr>
      <th scope="row">TX</th>
      <td>Provides an output UART to the DUT, this is connected to the jumpstarter console.</td>
    </tr>
    <tr>
      <th scope="row">RX</th>
      <td>Provides an input UART from the DUT, this is connected to the jumpstarter console.</td>
    </tr>
    <tr>
      <th scope="row">CTL_A</th>
      <td>Provides a digital input/output which can be controlled by the jumpstarter software.
          The current convention is to use this pin for forcing devices into flashing mode,
          but it can be used for any purpose.</td>
    </tr>
    <tr>
      <th scope="row">CTL_B</th>
      <td>Provides a digital input/output which can be controlled by the jumpstarter software.</td>
    </tr>
    <tr>
      <th scope="row">CTL_C</th>
      <td>Provides a digital input/output which can be controlled by the jumpstarter software.</td>
    </tr>
    <tr>
      <th scope="row">CTL_D</th>
      <td>Provides a digital input/output which can be controlled by the jumpstarter software.</td>
    </tr>
    <tr>
      <th scope="row">/RESET</th>
      <td>Provides a reset output signal which can be controlled from the jumpstarter
              software. It's active low open collector output, this means that it will output
              a '0' when the reset is asserted, and it will be in HiZ when the reset is not.</td>
    </tr>
  </tbody>
</table>
