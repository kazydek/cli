package deploy

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kyma-project/cli/internal/cli"
)

var (
	defaultSource         = "master"
	kymaProfiles          = []string{"evaluation", "production"}
	defaultWorkspacePath  = filepath.Join(".", "workspace")
	defaultComponentsFile = filepath.Join(defaultWorkspacePath, "installation", "resources", "components.yaml")
	defaultTLSCrt         = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURQVENDQWlXZ0F3SUJBZ0lSQVByWW0wbGhVdUdkeVNCTHo4d3g5VGd3RFFZSktvWklodmNOQVFFTEJRQXcKTURFVk1CTUdBMVVFQ2hNTVkyVnlkQzF0WVc1aFoyVnlNUmN3RlFZRFZRUURFdzVzYjJOaGJDMXJlVzFoTFdSbApkakFlRncweU1EQTNNamt3T1RJek5UTmFGdzB6TURBM01qY3dPVEl6TlROYU1EQXhGVEFUQmdOVkJBb1RER05sCmNuUXRiV0Z1WVdkbGNqRVhNQlVHQTFVRUF4TU9iRzlqWVd3dGEzbHRZUzFrWlhZd2dnRWlNQTBHQ1NxR1NJYjMKRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFDemE4VEV5UjIyTFRKN3A2aXg0M2E3WTVVblovRkNicGNOQkdEbQpxaDRiRGZLcjFvMm1CYldWdUhDbTVBdTBkeHZnbUdyd0tvZzJMY0N1bEd5UXVlK1JLQ0RIVFBJVjdqZEJwZHJhCkNZMXQrNjlJMkJWV0xiblFNVEZmOWw3Vy8yZFFFU0ExZHZQajhMZmlrcEQvUEQ5ekdHR0FQa2hlenVNRU80dUwKaUxXSloyYmpYK1dtaGZXb0lrOG5oak5YNVBFN2l4alMvNnB3QU56eXk2NW95NDJPaHNuYXlDR1grbmhFVk5SRApUejEraEMvdjJaOS9lRG1OdHdjT1hJSk4relZtUTJ4VHh2Sm0rbDUwYzlnenZTY3YzQXg0dUJsOTk3UnVlcUszCmdZMVRmVklFQ0FOTE9hb29jRG5kcW1FY1lBb25SeGJKK0M2U1RJYlhuUVAyMmYxQkFnTUJBQUdqVWpCUU1BNEcKQTFVZER3RUIvd1FFQXdJRm9EQVRCZ05WSFNVRUREQUtCZ2dyQmdFRkJRY0RBVEFNQmdOVkhSTUJBZjhFQWpBQQpNQnNHQTFVZEVRUVVNQktDRUNvdWJHOWpZV3d1YTNsdFlTNWtaWFl3RFFZSktvWklodmNOQVFFTEJRQURnZ0VCCkFBUnVOd0VadW1PK2h0dDBZSWpMN2VmelA3UjllK2U4NzJNVGJjSGtyQVhmT2hvQWF0bkw5cGhaTHhKbVNpa1IKY0tJYkJneDM3RG5ka2dPY3doNURTT2NrdHBsdk9sL2NwMHMwVmFWbjJ6UEk4Szk4L0R0bEU5bVAyMHRLbE90RwpaYWRhdkdrejhXbDFoRzhaNXdteXNJNWlEZHNpajVMUVJ6Rk04YmRGUUJiRGkxbzRvZWhIRTNXbjJjU3NTUFlDCkUxZTdsM00ySTdwQ3daT2lFMDY1THZEeEszWFExVFRMR2oxcy9hYzRNZUxCaXlEN29qb25MQmJNYXRiaVJCOUIKYlBlQS9OUlBaSHR4TDArQ2Nvb1JndmpBNEJMNEtYaFhxZHZzTFpiQWlZc0xTWk0yRHU0ZWZ1Q25SVUh1bW1xNQpVNnNOOUg4WXZxaWI4K3B1c0VpTUttND0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
	defaultTLSKey         = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb2dJQkFBS0NBUUVBczJ2RXhNa2R0aTB5ZTZlb3NlTjJ1Mk9WSjJmeFFtNlhEUVJnNXFvZUd3M3lxOWFOCnBnVzFsYmh3cHVRTHRIY2I0SmhxOENxSU5pM0FycFJza0xudmtTZ2d4MHp5RmU0M1FhWGEyZ21OYmZ1dlNOZ1YKVmkyNTBERXhYL1plMXY5blVCRWdOWGJ6NC9DMzRwS1EvencvY3hoaGdENUlYczdqQkR1TGk0aTFpV2RtNDEvbApwb1gxcUNKUEo0WXpWK1R4TzRzWTB2K3FjQURjOHN1dWFNdU5qb2JKMnNnaGwvcDRSRlRVUTA4OWZvUXY3OW1mCmYzZzVqYmNIRGx5Q1RmczFaa05zVThieVp2cGVkSFBZTTcwbkw5d01lTGdaZmZlMGJucWl0NEdOVTMxU0JBZ0QKU3ptcUtIQTUzYXBoSEdBS0owY1d5Zmd1a2t5RzE1MEQ5dG45UVFJREFRQUJBb0lCQUJwVmYvenVFOWxRU3UrUgpUUlpHNzM5VGYybllQTFhtYTI4eXJGSk90N3A2MHBwY0ZGQkEyRVVRWENCeXFqRWpwa2pSdGlobjViUW1CUGphCnVoQ0g2ZHloU2laV2FkWEVNQUlIcU5hRnZtZGRJSDROa1J3aisvak5yNVNKSWFSbXVqQXJRMUgxa3BockZXSkEKNXQwL1o0U3FHRzF0TnN3TGk1QnNlTy9TOGVvbnJ0Q3gzSmxuNXJYdUIzT1hSQnMyVGV6dDNRRlBEMEJDY2c3cgpBbEQrSDN6UjE0ZnBLaFVvb0J4S0VacmFHdmpwVURFeThSSy9FemxaVzBxMDB1b2NhMWR0c0s1V1YxblB2aHZmCjBONGRYaUxuOE5xY1k0c0RTMzdhMWhYV3VJWWpvRndZa0traFc0YS9LeWRKRm5acmlJaDB0ZU81Q0I1ZnpaVnQKWklOYndyMENnWUVBd0gzeksvRTdmeTVpd0tJQnA1M0YrUk9GYmo1a1Y3VUlkY0RIVjFveHhIM2psQzNZUzl0MQo3Wk9UUHJ6eGZ4VlB5TVhnOEQ1clJybkFVQk43cE5xeWxHc3FMOFA1dnZlbVNwOGNKU0REQWN4RFlqeEJLams5CldtOXZnTGpnaERSUFN1Um50QXNxQVVqcWhzNmhHUzQ4WUhMOVI2QlI5dmY2U2xWLzN1NWMvTXNDZ1lFQTdwM1UKRDBxcGNlM1liaiszZmppVWFjcTRGcG9rQmp1MTFVTGNvREUydmZFZUtEQldsM3BJaFNGaHYvbnVqaUg2WWJpWApuYmxKNVRlSnI5RzBGWEtwcHNLWW9vVHFkVDlKcFp2QWZGUzc2blZZaUJvMHR3VzhwMGVCS3QyaUFyejRYRmxUCnpRSnNOS1dsRzBzdGJmSzNqdUNzaWJjYTBUd09lbTdSdjdHV0dLTUNnWUJjZmFoVVd1c2RweW9vS1MvbVhEYisKQVZWQnJaVUZWNlVpLzJoSkhydC9FSVpEY3V2Vk56UW8zWm9Jc1R6UXRXcktxOW56VmVxeDV4cnkzd213SXExZwpCMFlVQVhTRlAvV1ZNWEtTbkhWVzdkRUs2S3pmSHZYTitIRjVSbHdLNmgrWGVyd2hsS093VGxyeVAyTEUrS1JtCks1cHJ5aXJZSWpzUGNKbXFncG9IbFFLQmdCVWVFcTVueFNjNERYZDBYQ0Rua1BycjNlN2lKVjRIMnNmTTZ3bWkKVVYzdUFPVTlvZXcxL2tVSjkwU3VNZGFTV3o1YXY5Qk5uYVNUamJQcHN5NVN2NERxcCtkNksrWEVmQmdUK0swSQpNcmxGT1ZpU09TZ1pjZUM4QzBwbjR2YXJFcS9abC9rRXhkN0M2aUhJUFhVRmpna3ZDUllIQm5DT0NCbjl4TUphClRSWlJBb0dBWS9QYSswMFo1MHYrUU40cVhlRHFrS2FLZU80OFUzOHUvQUJMeGFMNHkrZkJpUStuaXh5ZFUzOCsKYndBR3JtMzUvSU5VRTlkWE44d21XRUlXVUZ3YVR2dHY5NXBpcWNKL25QZkFiY2pDeU8wU3BJWCtUYnFRSkljbgpTVjlrKzhWUFNiRUJ5YXRKVTdIQ3FaNUNTWEZuUnRNanliaWNYYUFKSWtBQm4zVjJ3OFk9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="
)

//Options defines available options for the command
type Options struct {
	*cli.Options
	WorkspacePath  string
	ComponentsFile string
	OverridesFile  string
	Overrides      []string
	CancelTimeout  time.Duration
	QuitTimeout    time.Duration
	HelmTimeout    time.Duration
	WorkersCount   int
	Domain         string
	TLSCrt         string
	TLSKey         string
	Source         string
	Profile        string
}

//NewOptions creates options with default values
func NewOptions(o *cli.Options) *Options {
	return &Options{Options: o}
}

//profiles returns the currently supported profiles
func (o *Options) profiles() []string {
	return kymaProfiles
}

//defaultSource returns the default Kyma version
func (o *Options) defaultSource() string {
	return defaultSource
}

//defaultWorkspacePath returns the default path to the CLI workspace directory
func (o *Options) defaultWorkspacePath() string {
	return defaultWorkspacePath
}

//defaultComponentsFile returns the default path to the components list file
func (o *Options) defaultComponentsFile() string {
	return defaultComponentsFile
}

//defaultTLSCert returns the default path TLS certificate (base64 encoded)
func (o *Options) defaultTLSCrt() string {
	return defaultTLSCrt
}

//defaultComponentsFile returns the default TLS key (base64 encoded)
func (o *Options) defaultTLSKey() string {
	return defaultTLSKey
}

func (o *Options) supportedProfile(profile string) bool {
	for _, supportedProfile := range kymaProfiles {
		if supportedProfile == profile {
			return true
		}
	}
	return false
}

// validateFlags applies a sanity check on provided options
func (o *Options) validateFlags() error {
	// Overrides file is optional, but if provided it has to exist
	if o.OverridesFile != "" {
		if err := o.pathExists(o.OverridesFile, "Overrides file"); err != nil {
			return err
		}
	}
	if o.QuitTimeout < o.CancelTimeout {
		return fmt.Errorf("Quit timeout (%v) cannot be smaller than cancel timeout (%v)", o.QuitTimeout, o.CancelTimeout)
	}
	if o.Profile != "" && !o.supportedProfile(o.Profile) {
		return fmt.Errorf("Profile unknown or not supported. Supported profiles are: %s", strings.Join(o.profiles(), ", "))
	}
	if o.Domain != LocalKymaDevDomain && !o.tlsCertAndKeyProvided() {
		return fmt.Errorf("To use a custom domain name also a custom TLS certificate and TLS key has to be provided")
	}
	if (o.TLSKey != "" || o.TLSCrt != "") && !o.tlsCertAndKeyProvided() {
		return fmt.Errorf("To use a custom TLS certificate the TLS certificate and TLS key has to be provided")
	}
	return nil
}

func (o *Options) tlsCertAndKeyProvided() bool {
	return o.TLSCrt != "" && o.TLSKey != ""
}

func (o *Options) pathExists(path string, description string) error {
	if path == "" {
		return fmt.Errorf("%s is empty", description)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%s '%s' not found", description, path)
	}
	return nil
}