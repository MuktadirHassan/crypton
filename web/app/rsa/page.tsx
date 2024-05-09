'use client'
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { API_URL } from "@/config/consts";
import { useState } from "react";

async function Encryption(n: string, publicK: string, message: string) {
    const response = await fetch(API_URL + "/rsa/encrypt", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ n, public: publicK, message }),
    });

    if (!response.ok) {
        throw new Error(response.statusText);
    }

    return response.json();
}

async function Decryption(n: string, privateK: string, ciphertext: string) {
    const response = await fetch(API_URL + "/rsa/decrypt", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ n, private: privateK, ciphertext }),
    });
    return response.json();
}

async function KeyGeneration() {
    const response = await fetch(API_URL + "/rsa/keygen", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
    });

    if (!response.ok) {
        throw new Error(response.statusText);
    }

    return response.json();
}

export default function Cipher() {
    return (
        <div className="my-20">
            <h1 className="text-2xl mb-2">RSA Cipher</h1>
            <p className="mb-4">
                RSA is an asymmetric cryptographic algorithm. It uses a pair of keys to encrypt and decrypt data.
            </p>
            <KeyGenerator />
            <h1>
                Enter the secret key and text to encrypt or decrypt.
            </h1>
            <div className="flex justify-between">
                <CipherBox
                    type="encrypt"
                    fn={Encryption}
                />
                <CipherBox
                    type="decrypt"
                    fn={Decryption}
                />
            </div>
        </div>
    );
}

function KeyGenerator() {
    const { toast } = useToast();
    const [publicKey, setPublicKey] = useState("");
    const [privateKey, setPrivateKey] = useState("");
    const [n, setN] = useState(0);

    const handleGenerate = async () => {
        try {
            const response = await KeyGeneration();
            setPublicKey(response.public);
            setPrivateKey(response.private);
            setN(response.n);
        } catch (error: any) {
            console.log(error.message);
            toast({
                title: "Error",
                description: error.message,
            })
        }
    };

    return (
        <div className="flex flex-col gap-4 mb-20">
            <Button onClick={handleGenerate}>Generate Key</Button>
            <div className="flex gap-4 justify-between">
                <textarea
                    className="w-96 h-48 border border-gray-300 p-2"
                    value={publicKey}
                    readOnly
                    placeholder="Public Key"
                />
                <textarea
                    className="w-96 h-48 border border-gray-300 p-2"
                    value={privateKey}
                    readOnly
                    placeholder="Private Key"
                />
                <textarea
                    className="w-96 h-48 border border-gray-300 p-2"
                    value={n}
                    readOnly
                    placeholder="N"
                />
            </div>

        </div>
    );
}

function CipherBox({
    type,
    fn,
}: {
    type: "encrypt" | "decrypt",
    fn: any;
}) {
    const { toast } = useToast();
    const [text, setText] = useState("");
    const [key, setKey] = useState("");
    const [n, setN] = useState("");

    const [output, setOutput] = useState("");

    const handleSubmit = async (e: any) => {
        e.preventDefault();
        try {
            const response = await fn(n, key, text);
            console.log(response);
            if (response.ciphertext) {
                setOutput(response.ciphertext);
            }

            if (response.plaintext) {
                setOutput(response.plaintext);
            }
        } catch (error: any) {
            console.log(error.message);
            toast({
                title: "Error",
                description: error.message,
            })
        }

    };

    return (
        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
            <h1>{type === "encrypt" ? "Encryption" : "Decryption"}</h1>
            <textarea
                className="w-96 h-48 border border-gray-300 p-2"
                value={n}
                onChange={(e) => setN(e.target.value)}
                placeholder="N"
            />
            <textarea
                className="w-96 h-48 border border-gray-300 p-2"
                value={key}
                placeholder="Public Key"
                onChange={(e) => setKey(e.target.value)}
            />
            <textarea
                className="w-96 h-48 border border-gray-300 p-2"
                value={text}
                placeholder={type === "encrypt" ? "Enter plaintext" : "Enter ciphertext"}
                onChange={(e) => setText(e.target.value)}
            />
            <Button className="mt-2" type="submit">
                {type === "encrypt" ? "Encrypt" : "Decrypt"}
            </Button>


            <textarea
                className="w-96 h-48 border border-gray-300 p-2"
                value={output}
                readOnly
                placeholder="Output will be shown here."
            />
        </form>
    );
}
