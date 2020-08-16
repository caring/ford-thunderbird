# Add any output values to this file
output "service_url" {
  vale = module.service_dns_record.fqdn
}
