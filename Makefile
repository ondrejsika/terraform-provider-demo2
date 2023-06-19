build-dev:
	mkdir -p example/.terraform-plugin-cache/local/ondrejsika/demo2/0.0.1/darwin_arm64/
	go build && cp terraform-provider-demo2 example/.terraform-plugin-cache/local/ondrejsika/demo2/0.0.1/darwin_arm64/terraform-provider-demo2_v0.0.1
