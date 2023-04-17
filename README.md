# CoTNetwork DID Specification

## DID Scheme

The CoTNetwork DID scheme consists of the following parts:

* Identifier: `did`
* Identifier for the DID method: `cot`
* Method-specific identifier: `43*44(base58char)`

Examples of valid DIDs:

```
did:cot:F8otk18ArWDQrY1Q2XWkcZq3Te4VuuDWqn8mnc5t1KPa
did:cot:HSk6tA4V1uiSQhSyEdiUTU1ivoCfJMiEgPHE2raWzji
```

## DID Document

Example of CoTNetwork DID Document:

```json
{
  "@context":[
    "https://www.w3.org/ns/did/v1",
    "https://w3id.org/security/suites/bls12381-2020/v1"
    ],
    "id":"did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH","controller":"did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
    "verificationMethod":[{
      "id":"did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH#owner",
      "type":"Bls12381G2Key2020",
      "controller":"did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
      "publicKeyBase58":"nrYbjHKpmUHbDfwbco7Sn637t1YjoMVqzKjzSGwicLmfybMSYbyRMffGeUE3H9jNAjcMXrN5SK8Rc9KxKSCVCVRPcWH2A6EwrWrkt7tRzqwLwhD3uF3tvAA8UWg8yQzTYif"
    }
  ]
}
```

## DID Operations (CRUD)

### Create/Register

Users use [UniWa](https://ssis.cotnetwork.com) to create/register an entity in the CoT Network by submitting a Verifiable Presentation. The signature suite  **MUST** be [BbsBlsSignagure2020](https://w3c.github.io/vc-data-integrity/vocab/security/vocabulary.html#BbsBlsSignature2020), and the publickey **MUST** be [Bls12381G2Key2020]( https://w3c.github.io/vc-data-integrity/vocab/security/vocabulary.html#Bls12381G2Key2020). User **MUST** specifies a linkage which is created using the user's phone number. The phone number is encrypted by sha256 hash method. When login with phone number, the user can obtain a nonce via SMS text service at Uniwa register page.

```json
{
  "@context": [
    "https://www.w3.org/2018/credentials/v1",
    "https://w3id.org/security/bbs/v1"
  ],
  "proof": [
    {
      "created": "2023-02-11T10:42:03+08:00",
      "proofPurpose": "assertionMethod",
      "proofValue": "pwanOM4OITfB67El_p9kflxN9U0iKfQm5e9EavynB4DQqgGZKekmKGQfEF0pRbtmVtBaINGUhdwDd4ioOXRDcRU7CiV3qzlumIAS-TDYdlU8fMeab0Qq_Txiz80E7OnX7YBxtQbqCYm6kKxcIGGw_Q",
      "type": "BbsBlsSignature2020",
      "verificationMethod": "did:cot:6u3SCqoKfARwgbssjie1agpsoPitjKwkeFZxtJGb5BqY#owner"
    }
  ],
  "type": [
    "VerifiablePresentation"
  ],
  "verifiableCredential": [
    {
      "@context": [
        "https://ssis.cotnetwork.com/v1/schema/credentials/v1.1",
        "https://ssis.cotnetwork.com/v1/schema/bbs/v1.1",
        "https://ssis.cotnetwork.com/v1/schema/CoTNetworkIdentity/v1.0"
      ],
      "type": [
        "VerifiableCredential"
      ],
      "credentialSubject": {
        "linkage": "2e99758548972a8e8822ad47fa1017ff72f06f3ff6a016851f45c398732bc50c",
        "document": {
          "@context": [
            "https://www.w3.org/ns/did/v1",
            "https://w3id.org/security/suites/bls12381-2020/v1"
          ],
          "id": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
          "controller": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
          "verificationMethod": [
            {
              "id": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH#owner",
              "type": "Bls12381G2Key2020",
              "controller": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
              "publicKeyBase58": "nrYbjHKpmUHbDfwbco7Sn637t1YjoMVqzKjzSGwicLmfybMSYbyRMffGeUE3H9jNAjcMXrN5SK8Rc9KxKSCVCVRPcWH2A6EwrWrkt7tRzqwLwhD3uF3tvAA8UWg8yQzTYif"
            }
          ]
        },
        "nonce": "07890",
        "type": [
          "CoTNetworkIdentity"
        ]
      }
    }
  ]
}
```



The DID Document representing the entity is included in the body of the credentialSubject's document attribute. 

This transaction is approved and the entity is registered if all the following conditions are true:

1. If the `did` in the `proof.verificationMethod`  matches the `did` in one of `verifiableCredential.credentialSubject.document.verificationMethod` fields.
2. If the `proof` was successfully verified using `verifiableCredential.credentialSubject.document.verificationMethod`.
3. If the `verifiableCredential.credentialSubject.linkage` matches the `verifiableCredential.credentialSubject.nonce`.



### Read

Clients can read a DID document by sending a query request for a DID or a linkage.

For example, a query for a `did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH` or a linkage `2e99758548972a8e8822ad47fa1017ff72f06f3ff6a016851f45c398732bc50c` would return:

```json
{
  "@context": [
    "https://www.w3.org/ns/did/v1",
    "https://w3id.org/security/suites/bls12381-2020/v1"
  ],
  "id": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
  "controller": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
  "verificationMethod": [
    {
      "id": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH#owner",
      "type": "Bls12381G2Key2020",
      "controller": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
      "publicKeyBase58": "nrYbjHKpmUHbDfwbco7Sn637t1YjoMVqzKjzSGwicLmfybMSYbyRMffGeUE3H9jNAjcMXrN5SK8Rc9KxKSCVCVRPcWH2A6EwrWrkt7tRzqwLwhD3uF3tvAA8UWg8yQzTYif"
    }
  ]
}
```



### Update

Clients can update a DID document by submitting a Verifiable Presentation just like [Create/Register](#Create/Register).

```json
{
  "@context": [
    "https://www.w3.org/2018/credentials/v1",
    "https://w3id.org/security/bbs/v1"
  ],
  "proof": [
    {
      "created": "2023-02-11T10:42:03+08:00",
      "proofPurpose": "assertionMethod",
      "proofValue": "pwanOM4OITfB67El_p9kflxN9U0iKfQm5e9EavynB4DQqgGZKekmKGQfEF0pRbtmVtBaINGUhdwDd4ioOXRDcRU7CiV3qzlumIAS-TDYdlU8fMeab0Qq_Txiz80E7OnX7YBxtQbqCYm6kKxcIGGw_Q",
      "type": "BbsBlsSignature2020",
      "verificationMethod": "did:cot:6u3SCqoKfARwgbssjie1agpsoPitjKwkeFZxtJGb5BqY#owner"
    }
  ],
  "type": [
    "VerifiablePresentation"
  ],
  "verifiableCredential": [
    {
      "@context": [
        "https://ssis.cotnetwork.com/v1/schema/credentials/v1.1",
        "https://ssis.cotnetwork.com/v1/schema/bbs/v1.1",
        "https://ssis.cotnetwork.com/v1/schema/CoTNetworkIdentity/v1.0"
      ],
      "type": [
        "VerifiableCredential"
      ],
      "credentialSubject": {
        "linkage": "2e99758548972a8e8822ad47fa1017ff72f06f3ff6a016851f45c398732bc50c",
        "document": {
          "@context": [
            "https://www.w3.org/ns/did/v1",
            "https://w3id.org/security/suites/bls12381-2020/v1"
          ],
          "id": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
          "controller": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
          "verificationMethod": [
            {
              "id": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH#owner",
              "type": "Bls12381G2Key2020",
              "controller": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
              "publicKeyBase58": "nrYbjHKpmUHbDfwbco7Sn637t1YjoMVqzKjzSGwicLmfybMSYbyRMffGeUE3H9jNAjcMXrN5SK8Rc9KxKSCVCVRPcWH2A6EwrWrkt7tRzqwLwhD3uF3tvAA8UWg8yQzTYif"
            }
          ],
          [
            {
              "id": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH#signer",
              "type": "Bls12381G2Key2020",
              "controller": "did:cot:EhXSR8W1fLJnhaYQ3g8BLcszVsDorBs6NV7YfBe4rWgH",
              "publicKeyBase58": "nrYbjHKpmUHbDfwbco7Sn637t1YjoMVqzKjzSGwicLmfybMSYbyRMffGeUE3H9jNAjcMXrN5SK8Rc9KxKSCVCVRPcWH2A6EwrWrkt7tRzqwLwhD3uF3tvAA8UWg8yQzTYif"
            }
          ]
        },
        "nonce": "07890",
        "type": [
          "UniWaUserInformation"
        ]
      }
    }
  ]
}
```



This transaction is approved and the entity is registered if all the following conditions are true:

1. If the `did` in the `proof.verificationMethod`  matches the `did` in one of `verifiableCredential.credentialSubject.document.verificationMethod` fields.
2. If the `proof` was successfully verified using `verifiableCredential.credentialSubject.document.verificationMethod`.
3. If the `verifiableCredential.credentialSubject.linkage` matches the `verifiableCredential.credentialSubject.nonce`.

### Delete/Revoke

The DID document cannot be deleted/revoked from the chain.

### Status

The information about the status is in progess and will be updated once it is completed.



## Security Considerations

We only store DID documents on-chain. The user's DID document is generated locally, and the presentation method is used to prove its legitimacy based on the BBS+ signature method. Therefore, users can only operate on their own DID documents.

## Privacy Considerations

When the user registers, he needs to use SMS to verify the authenticity of his mobile phone number, but when generating the Presentation, the mobile phone number will be mapped to sha256 hash using the hash function. This process not only ensures the convenience of the user during use, but also prevents the leakage of the user's private information.

### References

1. Decentralized Identifiers (DIDs) v0.11 https://w3c-ccg.github.io/did-spec
2. Verifiable Credential v2.0 https://www.w3.org/2018/credentials/#verifiableCredential
3. Verifiable Presentation v2.0 https://www.w3.org/2018/credentials/#VerifiablePresentation
4. JSON-LD 1.0 - A JSON-based Serialization for Linked Data https://www.w3.org/TR/json-ld
5. BbsBlsSignature2020 https://w3c.github.io/vc-data-integrity/vocab/security/vocabulary.html#BbsBlsSignature2020
6. BbsBlsSignatureProof2020 https://w3c.github.io/vc-data-integrity/vocab/security/vocabulary.html#BbsBlsSignatureProof2020
7. Bls12381G2Key2020 https://w3c.github.io/vc-data-integrity/vocab/security/vocabulary.html#Bls12381G2Key2020

