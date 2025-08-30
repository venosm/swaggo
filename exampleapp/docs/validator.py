import openapi_spec_validator
import yaml

with open(
    "/home/milan/workspace_test/swag-master/exampleapp/docs/swagger.yaml", "r"
) as file:
    spec = yaml.safe_load(file)

try:
    print(openapi_spec_validator.OpenAPIV2SpecValidator(spec).is_valid())
    print(openapi_spec_validator.OpenAPIV30SpecValidator(spec).is_valid())
    print(openapi_spec_validator.OpenAPIV31SpecValidator(spec).is_valid())
    openapi_spec_validator.validate_spec(spec)
    print("✅ OpenAPI spec je validní!")
except Exception as e:
    print(f"❌ Chyby validace: {e}")
