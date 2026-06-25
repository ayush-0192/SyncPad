import "./globals.css";
import Sidebar from "@/component/Sidebar";

/*
export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
*/



export default function EditorLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <div className="flex h-screen bg-white gap-4">
        <div className="p-2 bg-black">
            <Sidebar />
        </div>
          <main className="flex-1">
            {children}
          </main>
        </div>
      </body>
    </html>
  );
}