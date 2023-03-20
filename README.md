# Salary Prediction dApp

This is a decentralized application (dApp) that predicts the salary of a job based on a dataset. The dApp is designed to run on the iExec platform and uses trusted execution 
environments (TEEs) to ensure that the computation is secure and private.


⚠️ This app was made in order to test the iExec stack, more than to really predict a possible salary 


## Dataset
The dataset used is encrypted (https://docs.iex.ec/for-developers/confidential-computing/sgx-encrypted-dataset) and the decryption key stored via the SMS component. Only the dApp has the authorization to access the decrypted dataset.

### Dataset format

The dataset is a simple CSV file with no headers:

```
job;number of years of experience, city, education, salary
```

Sample :

```
DEVOPS,2,PROVINCE,SELF-TAUGHT,33948
CLOUD_ARCHITECT,2,PARIS,ENGINEERING_SCHOOL,54450
```
I
Available data: 

 - JOBS : DEV_GOLANG,DEV_NODEJS,DEV_JAVA,CPP,C,ARCHITECT,DEV_PHP,DEV_PYTHON,FULL_STACK,FRONTEND,BACKEND,DEVOPS,CLOUD_ARCHITECT
 - CITY : PARIS, PROVINCE
 - EDUCATION :ENGINEERING_SCHOOL,COMPUTER_SCIENCE_SCHOOL,MASTER,BACHELOR


## Confidential Computing and TEE

In order to benefit from the computation confidentiality offered by Trusted Execution Environnements, I used the scon technology (https://docs.iex.ec/for-developers/confidential-computing/create-your-first-sgx-app)


## Run the dApp

The application expects in argument the name of the desired job, followed by the location, then the level of study and finally the number of years of experience.

For example for a nodejs developer job, in Paris with a bachelor level and 14 years of experience

``` shell
exec app run --args "DEV_NODEJS PARIS BACHELOR 14"  --tag tee,scone --dataset 0x4850b6e663A079F0022eC6D66EaE75FDe67d593f
```


