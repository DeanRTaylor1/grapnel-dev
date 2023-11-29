# Grapnel Solutions

## Overview

Grapnel Solutions is a streamlined web application focused on efficiency and portability. The entire application is embedded within a binary, ensuring a lightweight and robust deployment.

## Key Features

- **Embedded Application**: All resources and dependencies are bundled within a single binary.
- **Lightweight Container**: Thanks to a multi-stage build process, the Docker container size is optimized to just 27MB.

## Installation

```bash
docker build -t grapnel .
docker run -p 8080:8080 grapnel
```
