**SweetSSL** is a lightweight easy to use *reverse proxy* written in **Go** that provides FREE **Let's Encrypt** SSL certificates. It can run on **Linux** or as a **Windows** service.

It uses a mapping file to direct host names (`mysite.com`), a prefix (`/my-site`) or all traffic (`any`) to a backend server. The backend server can be a IP address, `HTTP`/`HTTPS` web server or can be a directory on disk (for static content).

The mapping file is watched on the file system and any changes are automatically added to the proxy without a restart.

### Download
You can download builds from the release page: [**Releases**](../../releases)

### How to use SSL with SweetSSL

1. Create a DNS entry for each host name pointing to your public IP address (A record) or create CNAME records pointing to an existing A record.
1. Update firewall settings so that both port 443 (SSL) and port 80 are accessible to the internet. Port 80 is required for Let's Encrypt to validate ownership of yoir domain.
1. Create an entry (one per line) for the host name entries created in Step 1 in the mapping.yml file (see examples below).
1. Start SweetSSL using the arguments below or install as a Windows service using the same command-line arguments.
1. Let SweetSSL do it's magic!

> Your email address is required when using the **Let's Encrypt** certificates. This is for **Let's Encrypt** to contact your about any issues.

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

```bash
Usage of sweetssl:
  -addr string
        Address to listen at (default ":https")
  -cache-dir string
        Path to directory to cache key and certificates (default Windows "C:\\ProgramData\sweetssl\cache", Linux "/var/cache/sweetssl")
  -email string
        Contact email address presented to letsencrypt CA
  -hostname string
        The default host name to be used with any and / prefix options
  -hsts
        Add Strict-Transport-Security header
  -http string
        Optional address to serve http-to-https redirects and ACME http-01 challenge responses (default ":http")
  -http-only
        Only use http
  -install
        Installs as a windows service
  -mapping string
        File with host/backend mapping (default "mapping.yml")
  -remove
        Removes the windows service
  -tls-skip-verify
        Skip verification of SSL certs for proxy targets
```

**SweetSSL** is a fork of the great [**leproxy**](https://github.com/artyom/leproxy) and uses the [**certmagic**](https://github.com/mholt/certmagic) library for **Let's Encrypt** support.
