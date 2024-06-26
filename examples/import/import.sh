#!/bin/sh
# Copyright (c) 2022 EPAM Systems, Inc.
# 
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.


hub api template create < import/template.json
hub api template init 10

# cd stack-k8s-aws && make deploy \
#   AWS_PROFILE=agilestacks-demo01 \
#   STATE_BUCKET=agilestacks.arkadi260.kubernetes.delivery STATE_REGION=eu-central-1 \
#   TF_VAR_name=arkadi267 TF_VAR_base_domain=arkadi260.kubernetes.delivery \
#   TF_VAR_aws_az=eu-central-1c \
#   TF_VAR_keypair=agilestacks

hub api instance create < import/instance.json
hub api secret stackInstance/arkadi267.arkadi260.kubernetes.delivery kubernetes.api.caCert certificate -     < import/arkadi267-ca-cert.pem
hub api secret stackInstance/arkadi267.arkadi260.kubernetes.delivery kubernetes.api.clientCert certificate - < import/arkadi267-client-cert.pem
hub api secret stackInstance/arkadi267.arkadi260.kubernetes.delivery kubernetes.api.clientKey privateKey -  < import/arkadi267-client-key.pem
hub api instance deploy 10
