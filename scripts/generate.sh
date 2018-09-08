docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
    -i /local/git.json \
    -l go \
    -o /local/out/go