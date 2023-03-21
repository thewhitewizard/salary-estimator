#!/bin/bash

# declare an image name
IMG_NAME="frederic1206/salary-estimator"

IMG_FROM=${IMG_NAME}:0.2
IMG_TO=${IMG_NAME}:0.2-tee-debug

SCONE_IMAGE="registry.scontain.com:5050/scone-production/iexec-sconify-image:5.7.5-v8"

# build the regular non-TEE image
docker build . -t ${IMG_FROM}

# run the sconifier to build the TEE image based on the non-TEE image
docker run -it \
            -v /var/run/docker.sock:/var/run/docker.sock \
            ${SCONE_IMAGE} \
            sconify_iexec \
              --cli=${SCONE_IMAGE} \
              --crosscompiler=${SCONE_IMAGE} \
              --name=salary-estimator \
              --from=${IMG_FROM} \
              --to=${IMG_TO} \
              --host-path=/etc/hosts \
              --host-path=/etc/resolv.conf \
              --binary="/app/salary-estimator" \
              --heap=1G \
              --dlopen=2 \
              --verbose \
              --no-color \
              --command="/app/salary-estimator" \
            && echo -e "\n------------------\n" \
            && echo "successfully built TEE docker image => ${IMG_TO}" \
            && echo "application mrenclave.fingerprint is $(docker run -it --rm -e SCONE_HASH=1 ${IMG_TO})"