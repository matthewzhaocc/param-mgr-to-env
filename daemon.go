// one that listens for a parameter and puts it in env
package main

import (
	// interact with OS environment
	"fmt"
	"os"
	"time"

	// interact with Systems Manager
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

)

var (
	Parameter string
	EnvName string
	AwsRegion string
)

// get that values that runs stuff
func init() {
	if len(os.Args) != 4 {
		fmt.Println("unable to parse arguments,\n please run it with the arguements param-mgr-to-env ssmVal envName regionName")
		return
	}
	
	args := os.Args[1:]
	Parameter = args[0]
	EnvName = args[1]
	AwsRegion = args[2]
}

// runs the main loop
func main() {
	sess := session.New(&aws.Config{
		Region: aws.String(AwsRegion),
	})
	svc := ssm.New(sess)
	// control loop
	for true {
		// god damn why is it in nanoseconds
		time.Sleep(10000000000)
		// query SSM
		res, err := svc.GetParameter(&ssm.GetParameterInput{
			Name: &Parameter,
		})
		// error checking
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// set it
		os.Setenv(EnvName, *res.Parameter.Value)
	}
}