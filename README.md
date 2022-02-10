# go-ini
INI Loader written in Go
Single threaded & simple

# Examples
* Read all params
    ```go
    func (app MyApp) onParam(name string, value string) bool {
        app.config.setParam(name, value)

        return true
    }

    func (app MyApp) Load() {
        reader := ini.NewReader()
        reader.ReadAll(app.dbConfig, app)
    }
    ```

# Requirements
* Go 1.17

# Tests
```bash
$go test
```
