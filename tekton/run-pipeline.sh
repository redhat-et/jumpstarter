#!/bin/sh

tkn pipeline start jumpstarter-orin-nx \
	--workspace name=checkout-files,volumeClaimTemplateFile=workspace-template.yaml \
       	--workspace name=image-input,claimName=lvm-rhel-image \
		--workspace name=image-files,volumeClaimTemplateFile=workspace-templates/image-workspace-template.yaml \
       	--workspace name=artifacts,volumeClaimTemplateFile=workspace-templates/workspace-template.yaml \
       	--showlog
