Use this plugin for publishing files and artifacts to Azure Storage. The plugin is power by [blobxfer](https://github.com/Azure/azure-batch-samples/tree/master/Python/Storage) - an upload tool from the Microsoft HFC team. The following parameters are required:

* `account_key` - Storage Account Key, required for authentication
* `storage_account` - Storage Account name
* `container` - The target storage container
* `source` - Location of folder to sync relative to the workspace root
* `target` - Remote resource location

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
publish:
  azure_storage:
    account_key: 123889asd89u8hsfdjh98128hh
    storage_account: my-storage-account
    container: my-storage-container
    source: folder/to/upload
    target: target/path
```
