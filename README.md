# did
A dead-simple simple library for generating prefixed IDs that provide context to consumers and creators alike.

## `did`s are:
- **Distinct**, as all unique identifiers should be
- **Definite**: "free of all ambiguity, uncertainty, or obscurity," thanks to prefixes
- **Dang...** way easier to use than UUIDs `;)`

## usage
`#TODO`

## features
### implemented
1. Create random did given a prefix
1. Create did from UUID or properly-formatted string
1. Methods for String, Length

### planned
1. more validations
1. built-in prefixes for common use-cases (e.g "users", "accounts")
1. More UUID interop: (e.g. comparison)
1. #Scan, #Value for DB reading + writing
1. performance / benchmarking
1. sorting?
1. other languages/ alphabets beyond English

### background
At the last company I worked for, I joined very early on when the product was a simple prototype / demo. I immediately saw an opportunity to simplify and standardize the database schema. I proposed common metadata fields for all records, decided on a standard approach to "soft-deleting" records (where necessary), and prioritized creating a standard approach for records IDs that would make records far easier to reason about than UUIDs. There was a stark contrast to the UUID-only approach we had versus what I was used to from working at Twilio: [What is a SID?](https://www.twilio.com/docs/glossary/what-is-a-sid). And I knew that other developer-friendly companies took similar approaches. For example, Stripe also has prefixed IDs: [Designing APIs for humans: Object IDs](https://dev.to/stripe/designing-apis-for-humans-object-ids-3o5a). It turns out that Github and many others were taking similar approaches: [Behind GitHub’s new authentication token formats](https://github.blog/2021-04-05-behind-githubs-new-authentication-token-formats/).