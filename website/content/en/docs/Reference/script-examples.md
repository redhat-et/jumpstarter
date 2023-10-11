---
title: Script examples
weight: 8
date: 2023-10-11
description: Jumpstarter scripts in action
---

{{% pageinfo %}}
The jumpstarter scripting language allows automation of the test process.
{{% /pageinfo %}}

The following example is a script that will deploy a RHEL image that
has been built for the device, it contains a set of kernel modules
that need testing, and an updated kernel in a local repository.

It can be ran as follows:

```bash
$ sudo jumpstarter run-script script-examples/orin-agx.yaml
```

```yaml
name: "Test Jetson kmods"
selector:
  - orin
expect-timeout: 60

steps:
  - comment: "Powering off and writing the image to disk"

  - power: off

  - set-disk-image:
      image: "rhel-guest-image.raw"

  - storage: attach

  - comment: "Booting up and waiting for login prompt"

  - power: on

  - expect:
      this: "login: "
      debug_escapes: true
      timeout: 250

  - send:
      this:
        - "root\n"
        - "redhat\n"

  - pause: 3

  - expect:
      this: "[root@localhost ~]#"
      debug_escapes: false

  - comment: "Updating kernel if necessary and installing the jetpack kmods"

  - send:
      echo: true
      this:
        - "sudo dnf update -y\n"

  - expect:
      timeout: 120
      echo: true
      debug_escapes: false
      this: "Complete"

  - expect:
      debug_escapes: false
      this: "[root@localhost ~]#"

  - send:
      echo: true
      this:
        - "sudo dnf install -y nvidia-jetpack-kmod\n"

  - expect:
      timeout: 120
      echo: true
      debug_escapes: false
      this: "Complete"

  - comment: "Rebooting to get latest kernel and kmods"

  - send:
      debug_escapes: false
      this:
        - "reboot\n"

  - expect:
      this: "login: "
      debug_escapes: false
      timeout: 300 # the kmod boot takes very long because of some issues with the crypto modules from nvidia

  - send:
      this:
        - "root\n"
        - "redhat\n"

  - send:
      echo: false # we dont want to capture any of the output so expect will catch it later
      this:
        - "\n"
        - "\n"

  - expect:
      debug_escapes: false
      echo: true
      this: "[root@localhost ~]#"

  - comment: "verifying that the kmods are loaded"

  - send:
      echo: false # we dont want to capture any of the output so expect will catch it
      this:
        - "lsmod | grep --color=never nv\n"

  - expect:
      echo: true
      this: "nvgpu"

  - comment: "Creating an inventory for this device and continuing with ansible"

  - write-ansible-inventory:
      filename: "inventory.yaml"
      ssh_key: ~/.ssh/id_rsa

  - local-shell:
      script: |
        ansible -m ping -i inventory.yaml all

cleanup:
  - comment: "Powering off and detaching the disk"
  - send:
      debug_escapes: false
      this:
        - "poweroff\n"
  - pause:  5

  - power: off

  - storage: detach

```

The output for this script would look as follows:

```
$ sudo ./jumpstarter set-disk-image visionfive2-00 rhel-guest-image.raw
💾 Writing disk image for visionfive2-00
🔍 Detecting USB storage device and connecting to host: <c^C
[gitlab-runner@localhost jumpstarter]$ ^C
[gitlab-runner@localhost jumpstarter]$ sudo ./jumpstarter run-script script-examples/orin-agx.yaml
⚙ Using device "orin-nx-00" with tags [orin-nx-00 orin orin-nx 16gb]

➤ Powering off and writing the image to disk
➤ Step ➤ power: "off"
[✓] done

➤ Step ➤ set-disk-image
🔍 Detecting USB storage device and connecting to host: done
📋 rhel-guest-image.raw -> /dev/disk/by-id/usb-SanDisk_Extreme_Pro_52A456790D93-0:0 offset 0x0:
💾 10240 MB copied 272.79 MB/s
[✓] done

➤ Step ➤ storage: "attach"
[✓] done


➤ Booting up and waiting for login prompt
➤ Step ➤ power: "on"
[✓] done

➤ Step ➤ expect: "login: "
040000 (0xe580)
[0000.414] I> RAM_CODE 0x4000401
[0000.420] I> RAM_CODE 0x4000401
[0000.423] I> Task: Load Page retirement list (0x500115dd)
[0000.429] I> Task: SDRAM params override (0x50012279)
[0000.433] I> Task: Save mem-bct info (0x5001542d)
[0000.438] I> Task: Carveout allocate (0x50015315)

...
...

......  Booting `Red Hat Enterprise Linux (5.14.0-362.6.1.el9_3.aarch64) 9.3
(Plow)'

EFI stub: Booting Linux Kernel...
EFI stub: Using DTB from configuration table
EFI stub: Exiting boot services...
[    0.000000] Booting Linux on physical CPU 0x0000000000 [0x410fd421]
[    0.000000] Linux version 5.14.0-362.6.1.el9_3.aarch64 (mockbuild@arm64-026.build.eng.bos.redhat.com) (gcc (GCC) 11.4.1 20230605 (Red Hat 11.4.1-2), GNU ld version 2.35.2-42.el9) #1 SMP PREEMPT_DYNAMIC Fri Sep 29 13:05:29 EDT 2023

...
...



localhost login:
➤ Step ➤ send


sent: root

root
Password:

sent: redhat



➤ Step ➤ pause: 3
[✓] done

➤ Step ➤ expect: "[root@localhost ~]#"
[root@localhost ~]#

➤ Updating kernel if necessary and installing the jetpack kmods
➤ Step ➤ send


sent: sudo dnf update -y

 sudo dnf update -y

<ESC>[?2004l
➤ Step ➤ expect: "Complete"
jetson-packages                                 3.1 MB/s | 519 kB     00:00
Dependencies resolved.
================================================================================
 Package               Arch      Version                Repository         Size
================================================================================
Installing:
 kernel                aarch64   5.14.0-362.8.1.el9_3   jetson-packages   5.1 M
 kernel-core           aarch64   5.14.0-362.8.1.el9_3   jetson-packages    18 M
 kernel-modules        aarch64   5.14.0-362.8.1.el9_3   jetson-packages    23 M
 kernel-modules-core   aarch64   5.14.0-362.8.1.el9_3   jetson-packages    26 M

Transaction Summary
================================================================================
Install  4 Packages

Total size: 73 M
Installed size: 110 M
Downloading Packages:
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                        1/1
  Installing       : kernel-modules-core-5.14.0-362.8.1.el9_3.aarch64       1/4
  Installing       : kernel-core-5.14.0-362.8.1.el9_3.aarch64               2/4
  Running scriptlet: kernel-core-5.14.0-362.8.1.el9_3.aarch64               2/4
  Installing       : kernel-modules-5.14.0-362.8.1.el9_3.aarch64            3/4
  Running scriptlet: kernel-modules-5.14.0-362.8.1.el9_3.aarch64            3/4
  Installing       : kernel-5.14.0-362.8.1.el9_3.aarch64                    4/4
  Running scriptlet: kernel-modules-core-5.14.0-362.8.1.el9_3.aarch64       4/4
  Running scriptlet: kernel-core-5.14.0-362.8.1.el9_3.aarch64               4/4
  Running scriptlet: kernel-modules-5.14.0-362.8.1.el9_3.aarch64            4/4
  Running scriptlet: kernel-5.14.0-362.8.1.el9_3.aarch64                    4/4
  Verifying        : kernel-5.14.0-362.8.1.el9_3.aarch64                    1/4
  Verifying        : kernel-core-5.14.0-362.8.1.el9_3.aarch64               2/4
  Verifying        : kernel-modules-5.14.0-362.8.1.el9_3.aarch64            3/4
  Verifying        : kernel-modules-core-5.14.0-362.8.1.el9_3.aarch64       4/4

Installed:
  kernel-5.14.0-362.8.1.el9_3.aarch64
  kernel-core-5.14.0-362.8.1.el9_3.aarch64
  kernel-modules-5.14.0-362.8.1.el9_3.aarch64
  kernel-modules-core-5.14.0-362.8.1.el9_3.aarch64

Complete
➤ Step ➤ expect: "[root@localhost ~]#"
!
[root@localhost ~]#
➤ Step ➤ send


sent: sudo dnf install -y nvidia-jetpack-kmod

 sudo dnf install -y nvidia-jetpack-kmod

<ESC>[?2004l
➤ Step ➤ expect: "Complete"
Last metadata expiration check: 0:01:02 ago on Wed Oct 11 09:39:12 2023.
Dependencies resolved.
==========================================================================================
 Package                   Arch     Version                         Repository        Size
==========================================================================================
Installing:
 nvidia-jetpack-kmod       aarch64  6.0.0_pre_ea_5.14.0_362-3.el9_3 jetson-packages   47 M
Installing dependencies:
 nvidia-jetpack-firmware   aarch64  6.0.0_pre_ea-3.el9              jetson-packages  659 k
 nvidia-jetpack-modprobe   aarch64  6.0.0_pre_ea-3.el9              jetson-packages   12 k

Transaction Summary
==========================================================================================
Install  3 Packages

Total size: 48 M
Installed size: 318 M
Downloading Packages:
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                        1/1
  Installing       : nvidia-jetpack-modprobe-6.0.0_pre_ea-3.el9.aarch64     1/3
  Installing       : nvidia-jetpack-firmware-6.0.0_pre_ea-3.el9.aarch64     2/3
  Installing       : nvidia-jetpack-kmod-6.0.0_pre_ea_5.14.0_362-3.el9_3.   3/3
  Running scriptlet: nvidia-jetpack-modprobe-6.0.0_pre_ea-3.el9.aarch64     3/3
  Running scriptlet: nvidia-jetpack-kmod-6.0.0_pre_ea_5.14.0_362-3.el9_3.   3/3
  Verifying        : nvidia-jetpack-firmware-6.0.0_pre_ea-3.el9.aarch64     1/3
  Verifying        : nvidia-jetpack-kmod-6.0.0_pre_ea_5.14.0_362-3.el9_3.   2/3
  Verifying        : nvidia-jetpack-modprobe-6.0.0_pre_ea-3.el9.aarch64     3/3

Installed:
  nvidia-jetpack-firmware-6.0.0_pre_ea-3.el9.aarch64
  nvidia-jetpack-kmod-6.0.0_pre_ea_5.14.0_362-3.el9_3.aarch64
  nvidia-jetpack-modprobe-6.0.0_pre_ea-3.el9.aarch64

Complete

➤ Rebooting to get latest kernel and kmods
➤ Step ➤ send


sent: reboot

!
reboot

➤ Step ➤ expect: "login: "
[root@localhost ~]# reboot
[root@localhost ~]#

Red Hat Enterprise Linux 9.3 (Plow)
Kernel 5.14.0-362.8.1.el9_3.aarch64 on an aarch64

Activate the web console with: systemctl enable --now cockpit.socket

localhost login:
➤ Step ➤ send


sent: root

root
Password:

sent: redhat



➤ Step ➤ send


sent:



sent:


➤ Step ➤ expect: "[root@localhost ~]#"


Last login: Wed Oct 11 09:39:04 on ttyTCU0
[root@localhost ~]#

➤ verifying that the kmods are loaded
➤ Step ➤ send


sent: lsmod | grep --color=never nv


➤ Step ➤ expect: "nvgpu"


<ESC>[?2004l
<ESC>[?2004h[root@localhost ~]#

<ESC>[?2004l
<ESC>[?2004h[root@localhost ~]# lsmod | grep --color=never nv

nvgpu[?2004l

➤ Creating an inventory for this device and continuing with ansible
➤ Step ➤ write-ansible-inventory

written : inventory.yaml

➤ Step ➤ local-shell
+ ansible -m ping -i inventory.yaml all
orin-nx-00 | SUCCESS => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python3"
    },
    "changed": false,
    "ping": "pong"
}
➤ Cleanup ➤ Test Jetson kmods

➤ Powering off and detaching the disk
➤ Step ➤ send


sent: poweroff

poweroff
[root@localhost ~]#
         Stopping Session 1 of User root...
         Stopping Session 3 of User root...


➤ Step ➤ pause: 5
[✓] done

➤ Step ➤ power: "off"
[✓] done

➤ Step ➤ storage: "detach"
[✓] done

```