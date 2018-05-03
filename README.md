# get-ssm-params

Get AWS SSM parameters

## Build

Clone this repository. Inside checkout:

```bash
# if using Go < 1.8
export GOPATH=/home/<your_user>/go

# fetch dependencies (aws sdk, as listed in main.go) into $GOPATH
go get -d

# build; env required if building on mac for linux
env GOOS=linux GOARCH=amd64 go build

# alternatively, if not cross compiling -- install into $GOPATH/bin using
env GIT_TERMINAL_PROMPT=1 go get -v github.com/projectThor/get-ssm-params
```

## Usage

From -help output:

```
USAGE: go-ssm-params [file.json].
It will try to read a valid flat JSON from stdin if no file passed.

Examples reading from stdin:
----------------------------
        $ cat file.json | go-ssm-params
        $ go-ssm-params < file.json
```

### Examples

```bash
# Fetch two params for given env/service
$ get-ssm-params -env PROD -service FOOBAR -params DB_HOST,DB_USER,DB_PASS
# this will retrieve parameters named PROD_FOOBAR_DB_HOST and PROD_FOOBAR_DB_USER,
# but output will be (env/service is stripped):
DB_HOST="example.com"
DB_USER="..."

# or, get specific/explicit params (i.e. not relying on env/service)
$ get-ssm-params -extraparams BLA_BLI,BLUB_BLUB
BLA_BLI="yoyo"
BLUB_BLUB="yumyum"

# alternatively, let get-ssm-params put parameters into env and exec a 'final' entrypoint
$ export SSM_ENV=FOO
$ export SSM_SERVICE=BAR
$ export SSM_PARAMS=BLURP,BLIP
# this will exec 'npm run', with desired/retrieved params in environment:
get-ssm-params npm run
```

If one of the parameters cannot be retrieved, `get-ssm-params` will `exit(1)`.
By default, `get-ssm-params` uses AWS region `eu-central-1`. To override,
use `-awsregion` command line option or define `SSM_AWS_REGION` environment variable.

## Example error messages

```
NoCredentialProviders: no valid providers in chain. Deprecated.
  For verbose messaging see aws.Config.CredentialsChainVerboseErrors
```

The above message indicates that the host has no policy attached.

```
AccessDeniedException: User: arn:aws:sts::123412341234:assumed-role/myec2Role/i-c7722c4d is not authorized to perform: ssm:GetParameters on resource: arn:aws:ssm:eu-west-1:123412341234:parameter/testval
  status code: 400, request id: df465e18-109d-11e7-bfc5-4f01fca110c2
```

The above message indicates that the given ec2 host has a policy attached,
but it lacks permission on requested parameters.

## Links

 - https://aws.amazon.com/blogs/compute/managing-secrets-for-amazon-ecs-applications-using-parameter-store-and-iam-roles-for-tasks/
 - http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-walk.html
 - https://godoc.org/github.com/aws/aws-sdk-go/service/ssm#GetParametersInput
 - https://godoc.org/github.com/aws/aws-sdk-go/service/ssm#example-SSM-GetParameters
 - http://docs.aws.amazon.com/cli/latest/reference/ssm/get-parameters.html
 - https://github.com/aws/amazon-ssm-agent/blob/master/agent/parameters/parameters.go
 - https://github.com/aws/amazon-ssm-agent/blob/master/agent/parameterstore/parameterstore.go
