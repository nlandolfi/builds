docker run -it --env-file ./env.list --entrypoint bash --privileged gcr.io/isito-prow/bazel:0.22
#docker run -it --env-file ./env.list --privileged gcr.io/isito-prow/bazel:0.20
# ./test-infra-1/jenkins/bootstrap.py --repo="github.com/nlandolfi/test-infra" --job=basic-build --clean --pull=master:ef2d87b7436973a948aa29e929808c671629b390,3:9bf50c3ea0e5a9d9fc587cfd5f8f0faca4446253
# ./test-infra-1/jenkins/bootstrap.py --repo="github.com/nlandolfi/pilot" --job=pilot-presubmit --clean --pull=master:11dc95967a0e11c723922e51e07cb1b9935973e0,1:b0303fa0946f91a51a027ee2b6209e1407b7dfe7 --root="${GOPATH}/src"

