# Slack ChatOps Practice

Simple slack app builds and deploys with [Github Actions] *repository-dispatch*.

[Github Actions]: https://github.com/features/actions


## Flow to deploy

1. Execute command like `/deploy <commit-sha>` in slack.
1. Validate inputs. If it was invalid, abort.
1. Request Github Actions *repository-dispatch*, and then it will build and deploy on the commit-sha.

