'use client'
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { API_URL } from "@/config/consts";
import { useState } from "react";

async function Encryption(text: string, key: string) {
    const response = await fetch(API_URL + "/multiplicative/encrypt", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ plaintext: text, key }),
    });

    if (!response.ok) {
        throw new Error(response.statusText);
    }

    return response.json();
}

async function Decryption(text: string, key: string) {
    const response = await fetch(API_URL + "/multiplicative/decrypt", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ ciphertext: text, key }),
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
    fn: (text: string, key: string) => Promise<any>;
}) {
    const { toast } = useToast();
    const [text, setText] = useState("");
    const [key, setKey] = useState("");

    const [output, setOutput] = useState("");

    const handleSubmit = async (e: any) => {
        e.preventDefault();
        try {
            const response = await fn(text, key);
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
                value={key}
                placeholder="secret key"
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
