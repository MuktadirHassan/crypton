"use client"

import * as React from "react"
import Link from "next/link"

import { cn } from "@/lib/utils"
import {
    NavigationMenu,
    NavigationMenuItem,
    NavigationMenuLink,
    NavigationMenuList,
    navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu"

const menuItems = [
    { href: "/", label: "Home" },
    { href: "/rsa", label: "RSA" },
    { href: "/additive", label: "Additive" },

]


export function Header() {
    return (
        <NavigationMenu className="py-2">
            <NavigationMenuList>
                {
                    menuItems.map((item) => (
                        <NavigationMenuItem key={item.href}>
                            <Link href={item.href} passHref>
                                <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                                    {item.label}
                                </NavigationMenuLink>
                            </Link>
                        </NavigationMenuItem>
                    ))
                }
            </NavigationMenuList>
        </NavigationMenu>
    )
}

