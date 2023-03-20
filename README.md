# Salary Estimator dApp

This is a decentralized application (dApp) that predicts the salary of a job based on a dataset. The dApp is designed to run on the iExec platform and uses trusted execution 
environments (TEEs) to ensure that the computation is secure and private.


⚠️ This app was made in order to test the iExec stack, it's a POC. 



## How to build, deploy and run from scratch


In the following steps you will need :
 - An account on the registry https://hub.docker.com/ to be able to push your image
 - An account on https://sconedocs.github.io/registry/ to retrieve an image to sconify your application

Once the 2 accounts are created and activated, don't forget to make a docker login for these 2 registries

In addition, you will need the iExec CLI. For that follow the official documentation here https://docs.iex.ec/for-developers/quick-start-for-developers and once the installation is done, initialize a wallet:

``` Shell
iexec wallet create
```


### Clone this repo and init 

``` Shell
git clone git@github.com:thewhitewizard/salary-estimator.git
cd salary-estimator
iexec init --skip-wallet
```


### Init the app


``` Shell
docker build -t <dockerusername>/salary-estimator:<tag> .
```

### Build the App 

This module is written in Golang. A devcontainer is provided with the project making the developpement easy, but it is only if you need to modify the application sources.

**Step 1: Build and push the docker image**

``` Shell
docker build -t <dockerusername>/salary-estimator:<tag> .
```

In order to benefit from the computation confidentiality offered by Trusted Execution Environnements, we will use the scon technology (https://docs.iex.ec/for-developers/confidential-computing/create-your-first-sgx-app)

Modified line 4 of the sconfiy.sh script to indicate your user name on the docker hub, and line 6 to indicate your tag 

After, you can execute 
``` Shell
./sconify.sh    
```

Now you can push the app :
``` Shell
docker push <dockerusername>/salary-estimator:tee-debug
```


### Configuration and deploy

``` Shell
iexec app init --tee
iexec storage init --chain bellecour --tee-framework scone
```

according to the documentation , you have to get the fingerprint of your image and the fingerprint of the enclave


``` Shell
# mrenclave fingerprint
docker run -it --rm -e SCONE_HASH=1 <dockerusername>/salary-estimator:tee-debug 

#docker image fingerprint
docker pull  <dockerusername>/salary-estimator:tee-debug | grep "Digest: sha256:" | sed 's/.*sha256:/0x/' 
```

Edit iexec.jon and update values for *multiaddr, checksum and fingerprint*

Great, now you are ready to deploy
``` Shell
iexec app deploy --chain bellecour
```





### Dataset

The dataset used will be encrypted (https://docs.iex.ec/for-developers/confidential-computing/sgx-encrypted-dataset) and the decryption key stored via the SMS component. Only the dApp has the authorization to access the decrypted dataset.


The dataset is a simple CSV file with no headers and comma separator:

```
job;number of years of experience, city, education, salary
```

It is important to respect this format if you do not want to modify the sources

Sample :

```
DEVOPS,2,PROVINCE,SELF-TAUGHT,33948
CLOUD_ARCHITECT,2,PARIS,ENGINEERING_SCHOOL,54450
```

To make the dataset available to the dApp, follow the steps described here https://docs.iex.ec/for-developers/confidential-computing/sgx-encrypted-dataset



### Run this dApp

The application expects in argument the name of the desired job, followed by the location, then the level of study and finally the number of years of experience.

For example for a nodejs developer job, in Paris with a bachelor level and 14 years of experience

``` shell
exec app run 0x9535F5F413C764e76F2a1cf0f4e1526508B947BA --args "DEV_NODEJS PARIS BACHELOR 14"  --tag tee,scone --dataset 0x4850b6e663A079F0022eC6D66EaE75FDe67d593f
```
