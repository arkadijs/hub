#!/bin/sh
# Copyright (c) 2022 EPAM Systems, Inc.
# 
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.


hub api import eks arkadi352 -e Dev -w \
    --eks-endpoint 80A1E4C2EA4367436CA445DAA313204B.yl4.us-east-1.eks.amazonaws.com < examples/eks/arkadi352-ca-cert.pem

hub api template create < examples/eks/template.json
hub api template init 10

hub api instance create < examples/eks/instance.json
hub api secret stackInstance/arkadi352.arkadi351.kubernetes.delivery kubernetes.api.caCert certificate - < examples/eks/arkadi352-ca-cert.pem
hub api instance deploy 10
