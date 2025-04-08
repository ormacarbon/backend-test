export default function Footer() {
    return (
        <footer className="bg-white shadow-sm mt-auto">
            <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="py-4 text-center text-gray-600">
                    <p>© {new Date().getFullYear()} Cesar Brancalhão. All rights reserved.</p>
                </div>
            </div>
        </footer>
    );
}
