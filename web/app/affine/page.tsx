'use client'
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { API_URL } from "@/config/consts";
import { useState } from "react";

async function Encryption(text: string, key1: string, key2: string) {
    const response = await fetch(API_URL + "/affine/encrypt", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ plaintext: text, key1: Number(key1), key2: Number(key2) }),
    });

    if (!response.ok) {
        throw new Error(response.statusText);
    }

    return response.json();
}

async function Decryption(text: string, key1: string, key2: string) {
    const response = await fetch(API_URL + "/affine/decrypt", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ ciphertext: text, key1: Number(key1), key2: Number(key2) }),
    });
    return response.json();
}

export default function Cipher() {
    return (
        <div className="my-20">
            <h1 className="text-2xl mb-2">Affine Cipher</h1>
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

function CipherBox({
    type,
    fn,
}: {
    type: "encrypt" | "decrypt",
    fn: any;
}) {
    const { toast } = useToast();
    const [text, setText] = useState("");
    const [key1, setKey1] = useState(""); // key1 for multiplicative
    const [key2, setKey2] = useState(""); // key2 for additive

    const [output, setOutput] = useState("");

    const handleSubmit = async (e: any) => {
        e.preventDefault();
        try {
            const response = await fn(text, key1, key2);
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
            <input
                type="text"
                className="w-96 h-10 border border-gray-300 p-2"
                value={key1}
                placeholder="Enter key1"
                onChange={(e) => setKey1(e.target.value)}
            />
            <input type="text"
                className="w-96 h-10 border border-gray-300 p-2"
                value={key2}
                placeholder="Enter key2"
                onChange={(e) => setKey2(e.target.value)}

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
