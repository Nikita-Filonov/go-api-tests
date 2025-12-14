package http

import (
	"fmt"

	"github.com/Nikita-Filonov/axiom"
	"github.com/go-resty/resty/v2"
)

func onErrorHook(cfg *axiom.Config) func(req *resty.Request, err error) {
	return func(req *resty.Request, err error) {
		cfg.Step(fmt.Sprintf("Request error for %s %s", req.Method, req.URL), func() {
			if len(req.Header) > 0 {
				artefact, err := axiom.NewJSONArtefact("Request Headers", req.Header)
				if err == nil {
					cfg.Artefact(artefact)
				}
			}

			if req.Body != nil {
				artefact, err := axiom.NewJSONArtefact("Request Body", req.Body)
				if err == nil {
					cfg.Artefact(artefact)
				}
			}

			cfg.Artefact(
				axiom.NewTextArtefact("Error", err.Error()),
			)

			cfg.Log(
				axiom.NewErrorLog(
					fmt.Sprintf("request failed: %s %s", req.Method, req.URL),
				),
			)
		})
	}
}

func onBeforeRequestHook(cfg *axiom.Config) func(_ *resty.Client, req *resty.Request) error {
	return func(_ *resty.Client, req *resty.Request) error {
		cfg.Step(fmt.Sprintf("Send %s request to %s", req.Method, req.URL), func() {
			if len(req.Header) > 0 {
				artefact, err := axiom.NewJSONArtefact("Request Headers", req.Header)
				if err == nil {
					cfg.Artefact(artefact)
				}
			}

			if req.Body != nil {
				artefact, err := axiom.NewJSONArtefact("Request Body", req.Body)
				if err == nil {
					cfg.Artefact(artefact)
				}
			}
		})

		return nil
	}
}

func onAfterResponseHook(cfg *axiom.Config) func(_ *resty.Client, resp *resty.Response) error {
	return func(_ *resty.Client, resp *resty.Response) error {
		cfg.Step(fmt.Sprintf("Response: %s %s", resp.Request.Method, resp.Request.URL), func() {
			cfg.Artefact(
				axiom.NewTextArtefact("Response Status", resp.Status()),
			)

			cfg.Artefact(
				axiom.NewTextArtefact("Response Body", string(resp.Body())),
			)
		})

		return nil
	}
}
