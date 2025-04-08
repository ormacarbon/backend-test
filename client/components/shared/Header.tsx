import Link from 'next/link';

export default function Header() {
    return (
        <header className="bg-white shadow-sm">
            <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex justify-between h-16 items-center">
                    <Link href="/" className="font-bold text-xl text-blue-600">
                        VBIOS
                    </Link>
                    <nav>
                        <ul className="flex space-x-4">
                            <li>
                                <Link href="/" className="font-semibold text-gray-600 hover:text-blue-600">
                                    Leaderboard
                                </Link>
                            </li>
                            <li>
                                <Link href="/register" className="font-semibold text-gray-600 hover:text-blue-600">
                                    Register
                                </Link>
                            </li>
                        </ul>
                    </nav>
                </div>
            </div>
        </header>
    );
}
