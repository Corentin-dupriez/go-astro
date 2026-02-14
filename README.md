
# Planets Observer

**Planets Observer** is a Go application for calculating the positions of celestial
bodies (starting with the Sun and Earth) using astronomical algorithms. It
provides functions to compute mean longitude, mean anomaly, equation of center, true
longitude, true anomaly, and distance for a given Julian century (T).

---

## Features

- Calculate the **mean longitude** and **mean anomaly** of the Sun and Earth.
- Compute the **equation of center**, **true longitude**, **true anomaly**, and
**distance** from the Sun.
- Convert between civil dates and Julian dates.
- Modular and reusable Go package for astronomical calculations.

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Corentin-dupriez/planets_observer.git
   cd planets_observer
2. Ensure you have Go installed (version 1.21+ recommended):

    ```bash
    go version
3. Install dependencies
  
    ```bash
      go mod tidy

## Usage

### As a library

Import the astro package in your project

     ```go
       import "planets_observer/astro"

## Roadmap

- Add support for other planets
- Add ASCII visualisation
