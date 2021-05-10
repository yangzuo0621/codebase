package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	authorizer, err := auth.NewAuthorizerFromCLIWithResource(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		panic(err)
	}
	rclient := containerregistry.NewRegistriesClient("4be8920b-2978-43d7-ab14-04d8549c1d05")
	rclient.Authorizer = authorizer
	containerRegistryCredential, err := rclient.ListCredentials(context.Background(), "rp-common-e2e", "aksdeploymente2e")
	if err != nil {
		panic(err)
	}
	password := *containerRegistryCredential.Passwords

	c, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = c.RegistryLogin(ctx, types.AuthConfig{
		Username:      *containerRegistryCredential.Username,
		Password:      *password[0].Value,
		ServerAddress: "aksdeploymente2e.azurecr.io",
	})
	if err != nil {
		panic(err)
	}

	authConfig := types.AuthConfig{
		Username: *containerRegistryCredential.Username,
		Password: *password[0].Value,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	out, err := c.ImagePush(ctx, "aksdeploymente2e.azurecr.io/acs/overlaymgr:test", types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		panic(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
	if err != nil {
		panic(err)
	}

}
