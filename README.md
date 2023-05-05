# did
A dead-simple simple library for generating prefixed IDs that provide context to consumers and creators alike.

## `did`s are:
- **Distinct**, as all _unique_ identifiers are
- **Definite**: "free of all ambiguity, uncertainty, or obscurity," thanks to prefixes
- **Dang...** way easier to use than UUIDs `;)`

## usage
### creating a new, random did 
```
# a valid prefix is required
did, _ := did.New("us")
```

### creating a did from a UUID
```
u := uuid.New()
# a valid prefix is required
did, _ := did.DidFromUuid(u, "us")
```

### creating a did from a valid did string
```
did, _ := did.DidFromString("us-526cac357e74429beb4f2ecca56c571f")
```

## using a DidFactory
Instead of providing a prefix every time, you can use a factory to generate dids. Using a DidFactory is also the easiest way to use a separator other than the default (`-`).
```
df, _ := did.NewDidFactory("tot", "_")
d, _ := df.NewDid()
d.String() // tot_580a6ae69d3643289d83344c5925818c
```

## features
### implemented
1. Create random did given a prefix
1. Create did from UUID or properly-formatted string
1. Methods for String, Length
1. Validations for prefixes (alpha 2 or 3 characters)

### planned
1. built-in prefixes for common use-cases (e.g "users", "accounts")
1. More UUID interop: (e.g. comparison)
1. #Scan, #Value for DB reading + writing
1. performance / benchmarking
1. sorting?
1. other languages/ alphabets beyond English

### background
Prefixed IDs provide context that makes data from APIs or user interfaces far easier to reason about. If you've ever needed to copy or paste an API key, token, or even your account ID, you have likely mixed up one ID with another at some point. To be fair, this is a minor annoyance... but there is a better way. I was first exposed to prefixed IDs when working at Twilio: ([What is a SID?](https://www.twilio.com/docs/glossary/what-is-a-sid)), and the approach just _made sense_. Using the API, users often had to reference their account ID (AC prefix). Or if they wanted to get the status of an SMS (SM prefix), they needed that. Or a phone number (PN prefix). And so on. This would be far more cumbersome with UUID-heavy records or any other approach that ignores the user (developer) experience. Of course, Twilio is not the only company to do this. For example, Stripe also uses prefixed IDs: [Designing APIs for humans: Object IDs](https://dev.to/stripe/designing-apis-for-humans-object-ids-3o5a). And, as it turns out, Github and many others are doing the same: [Behind GitHub’s new authentication token formats](https://github.blog/2021-04-05-behind-githubs-new-authentication-token-formats/). 

### migration strategy
Moving tables from UUIDs or another type of unique ID will typically require adding a new column and "dual-writing" both UUIDs and dids. It's typically best to name the new column `id`, since dids are just one implementation of unique IDs. Start with nullable `id` columns, then change them to non-nullable _after_ completing backfills on your "pre-did" records. Update your foreign keys to reference dids, then... drop all the old columns! Most importantly, decide on an ID strategy early on to keep the migration simple. Or, better yet, start with prefixed IDs from the onset.