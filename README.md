# term-deposit-calculator

This Go program provides the final balance by taking below inputs:
- Start deposit amount (e.g. $10,000)
- Interest rate (e.g. 1.10%)
- Investment term (e.g. 3 years)
- Interest paid (monthly, quarterly, annually, at maturity)

## How to run the app

1. Ensure Go is installed.
2. Clone the repo.
3. Run:

```
make build
make run
```

## How to run the tests

```
make test
```

## How to run the linting

```
make lint
```

## Code considerations
1. Zerolog library is used for structured logging
2. The final balance is rounded to the nearest integer to match the Bendigo Bank calculator output.
3. Input validation for InterestPaid is enforced via enum allowed values (monthly, quarterly, annually, at-maturity).
4. The code implementation is kept very simple and clear.
5. The application accepts configuration via CLI flags or environment variables(as a fallback) defined in config.go.
6. Use Makefile run command to change values and run ```make run```

```
go run cmd/main.go \
  --start-deposit 10000 \
  --interest-rate 5.5 \
  --investment-term 3 \
  --interest-paid monthly
  ```

6. Compound interest is calculated using the formulae in https://moneysmart.gov.au/saving/compound-interest
7. If interest paid is at maturity, then a simple interest can be calculated using 
```
FinalBalance = P(1 + rt) where P is principal amount, r is Interest rate and t is investment term
```
