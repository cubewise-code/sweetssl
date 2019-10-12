**SweetSSL** is a lightweight easy to use *reverse proxy* written in **Go** that provides easy to use and FREE **Let's Encrypt** SSL certificates. It can run on **Linux** or as a **Windows** service.

It uses a mapping file to direct host names (`mysite.com`), a prefix (`/my-site`) or all traffic (`any`) to a backend server. The backend server can be a IP address, `HTTP`/`HTTPS` or can be a directory on disk (for static content).

The mapping file is watched on the file system and any changes are automatically added to the proxy without a restart.

Your email address is required when using the **Let's Encrypt** certificates. This is for **Let's Encrypt** to contact your about any issues.

Run `HTTPS` with default `mapping.yml` file:

	sweetssl -email youremail@yourdomain.com

Run `HTTPS` with custom mapping path:

	sweetssl -email youremail@yourdomain.com -mapping "otherfile.yml"

Run `HTTPS` and allow self signed certificates on backend servers:

	sweetssl -email youremail@yourdomain.com -tls-skip-verify

Run `HTTP` with default mapping:

	sweetssl -http-only

Install as a Windows service:

	sweetssl -install -email youremail@yourdomain.com

Get GoLang Source:

	go get github.com/cubewise-code/sweetssl

Build Go source:

	go build -o sweetssl.exe


`mapping.yml` contains host-to-backend mapping:

Example:

```yaml
   # Examples
   subdomain1.example.com: 127.0.0.1:8080
   uploads.example.com: https://uploads-bucket.s3.amazonaws.com
   static.linux.com: /var/www/
   static.windows.com: C:\Temp\
   /prod: http://prodserver/api/v1
   /dev: http://devserver/api/v1
   any: C:\Temp\
   any: https://localhost:8883/api/v1
```

**SweetSSL** is a fork of the great [**leproxy**](https://github.com/artyom/leproxy) and uses the [**certmagic**](https://github.com/mholt/certmagic) library for **Let's Encrypt** support.
