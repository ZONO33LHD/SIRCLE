import type { Metadata } from 'next';
import { Noto_Sans_JP } from 'next/font/google';
import Header from '@/components/Header';
import SideBar from '@/components/SideBar';
import '@/app/globals.css';
import { FinancialProvider } from '@/components/FinancialContext';
import ApolloWrapper from '@/lib/apollo-wrapper';

const noto = Noto_Sans_JP({ subsets: ['latin'], display: 'swap' });

export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={noto.className}>
        <ApolloWrapper>
          <FinancialProvider>
            <div className="flex h-screen flex-col">
              <Header />
              <div className="flex flex-1 overflow-hidden">
                <SideBar />
                <main className="flex flex-1 flex-col overflow-y-auto bg-gray-100">
                  <div className="flex-1">{children}</div>
                  <div className="h-16"></div>
                </main>
              </div>
            </div>
          </FinancialProvider>
        </ApolloWrapper>
      </body>
    </html>
  );
}
