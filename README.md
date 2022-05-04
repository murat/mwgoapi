# Merriam Webster Golang API Wrapper

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/eaa91a77066b494b8c357992f12a979b)](https://www.codacy.com/gh/murat/mwgoapi/dashboard?utm_source=github.com&utm_medium=referral&utm_content=murat/mwgoapi&utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/eaa91a77066b494b8c357992f12a979b)](https://www.codacy.com/gh/murat/mwgoapi/dashboard?utm_source=github.com&utm_medium=referral&utm_content=murat/mwgoapi&utm_campaign=Badge_Coverage)

## Usage examples

```golang
func main() {
  c := mwgoapi.NewClient(&http.Client{}, mwgoapi.BaseURL, "YOUR_API_KEY")
	r, err := c.Get("hello")
	if err != nil {
		fmt.Printf("could not get response, err: %v\n", err)
		os.Exit(1)
	}

	var collegiateResponse []mwgoapi.Collegiate
	if err := c.UnmarshalResponse(r, &collegiateResponse); err != nil {
		fmt.Printf("could not unmarshal response, err: %v\n", err)
		os.Exit(1)
	}

  fmt.Println(collegiateResponse)
}
```

-   [murat/dicterm](https://github.com/murat/dicterm)
