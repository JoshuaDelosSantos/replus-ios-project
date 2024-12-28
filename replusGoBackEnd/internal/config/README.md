# Config Layer Documentation

## Files and Functions
### config.go
This file is the core of the config package, containing the following components:
**type Config**
A struct that encapsulates all essential configuration parameters, including:
- DBHost: Database host address.
- DBPort: Database port.
- DBUser: Database username.
- DBPassword: Database password.
- DBName: Database name.
- AppPort: Application’s listening port.

**LoadConfig()**
This function initializes the Config struct by loading environment variables. Key features include:
- Attempts to load variables from a .env file using the godotenv package.
- Falls back to system environment variables if the .env file is not found.
- Assigns default values for each configuration parameter if the corresponding variable is not set.

**Example Usage:**
```
config := config.LoadConfig()
fmt.Println("Database Host:", config.DBHost)
```

**getEnv(key, defaultValue string) string**
A utility function for retrieving environment variables with a fallback mechanism.
- If the specified environment variable (key) exists, its value is returned.
- Otherwise, the provided defaultValue is returned.

### Purpose
The config directory serves as a single source of truth for application settings, enabling:
- Decoupling of configuration details from the codebase.
- Easy management of environment-specific variables.
- Use of default values for robust fallback mechanisms.