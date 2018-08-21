[![pipeline status](https://gitlab.com/mantris/corekit/badges/master/pipeline.svg)](https://gitlab.com/mantris/corekit/commits/master)
[![pnpm badge](https://img.shields.io/badge/package%20manager-pnpm-1abc9c.svg)](https://pnpm.js.org/)
[![JavaScript Style Guide](https://img.shields.io/badge/code_style-standard-brightgreen.svg)](https://standardjs.com)

# Corekit - Mantris' core systems

This is a monorepo.

It SECRET_KEYcontains all projects related to Mantris' core business as well as all **DevOps** stuff in order to make them work together.
We understand there're [upsides and downsides](https://github.com/IcaliaLabs/guides/wiki/Monolithic-vs-Micro-Repos) of using monorepos, so here's a few things to keep in mind:

**What do we need?**
- We need projects' applications to be make requests to each other, whenever necessary;
- We need all resources necessary to develop and deploy our projects versioned within this repository;
- We need to keep it simple so it would be easy for new developers/sysadmins to get onboard. Let's make it **APnPAP** (As Plug-n-Play As Possible);

**What would we love to have?**
- We would love each pull request to spin up a preview environment;
- We would love to locally run only a given project when there's no need to run the whole ecosystem;

**What should we be extra careful with?**
- We should avoid slow git performance;
- We shouldn't consume too much resources (CPU/memory) running the whole ecosystem locally;
