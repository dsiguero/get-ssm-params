package main

import (
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/ssm"
)

func SSMGet(keys []*string) map[string]string {
  sess := session.Must(session.NewSession())
  sessConfig := &aws.Config{}
  svc := ssm.New(sess, sessConfig)

  params := &ssm.GetParametersInput {
    Names: keys,
    WithDecryption: aws.Bool(true),
  }

  resp, err := svc.GetParameters(params)

  // TODO: Verify if there's any parameter in resp.InvalidParameters

  if err != nil {
    exitPanic(err)
  }

  var parameterMap map[string]string
  parameterMap = make(map[string]string)

  for _, v := range resp.Parameters {
    parameterMap[string(*v.Name)] = string(*v.Value)
  }

  return parameterMap
}
