# Transaction Services

This repository contains a series of cloud functions that are designed to help
manage transactions. This is implemented in Go and using Google Cloud Platform.

This is ultimately a logging system for transactions. It takes alerts which 
financial institutions send and put them into big query. 

From the outset, this isn't written to be helpful to anyone. It's more of a pet
project that I have created to: 
1. Learn Go lang.
2. Use cloud functions and microservices architecture for a problem.
3. Just help me with day to day logging of transactions.

That being said, if you come across this I hope you find it useful. I have made
it as "componetised" as possible. However, I'm a beginner to Go so probably not
using the best design pattern (or any design pattern) out there.

> If you're already using Google Cloud Platform, this should be plug and play.

# Getting Started

1. Simply clone down the repository
2. Deploy the cloud functions using gcloud functions deploy ...

# Available Functions

## GetTransaction

This is a simple function which extracts entities from a typical text alerts
that your financial institution may send you.

For example, you may receive the following.

```text
You spent $23.35 at McDonalds
```

Invoking this function will return you a Json as below
```text
{
    "Location":"McDonalds",
    "Amount":"$23.35",
    "NumericAmount":23.35,
    "NotifiedTime":"2019-10-07T10:03:42.868374665Z"
}
```

>*Underneath, this uses google cloud NLP api to do the entity extraction.*
>
>Associated charges may apply.
