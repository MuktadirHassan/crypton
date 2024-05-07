import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Header } from "@/components/custom/Header";
import { cn } from "@/lib/utils";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Crypton",
  description: "Crypton is a implementation of cryptographic algorithms in Go.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={cn(
        inter.className,
        "container",
        "mx-auto",
      )}>
        <Header />
        {children}
      </body>
    </html>
  );
}
