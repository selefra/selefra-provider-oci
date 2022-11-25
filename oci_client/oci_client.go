package oci_client

type OciClient struct {
	Compartment string
	Zone        string
	Region      string
	OciConfig   *OciConfig
}

func (x *OciClient) CopyWithRegion(region string) *OciClient {
	return &OciClient{
		Region:    region,
		OciConfig: x.OciConfig,
	}
}

func (x *OciClient) CopyWithCompartment(compartment string) *OciClient {
	return &OciClient{
		Compartment: compartment,
		OciConfig:   x.OciConfig,
	}
}

func (x *OciClient) CopyWithZone(zone string) *OciClient {
	return &OciClient{
		Zone:      zone,
		OciConfig: x.OciConfig,
	}
}

func (x *OciClient) CopyWithAll(region string, compartment string, zone string) *OciClient {
	return &OciClient{
		Zone:        zone,
		Compartment: compartment,
		Region:      region,
		OciConfig:   x.OciConfig,
	}
}
