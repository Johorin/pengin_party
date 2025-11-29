// route page.tsが呼ばれる前に実行されるミドルウェア

import "@/infrastructure/firebase.ts"
import { NextResponse, NextRequest } from 'next/server'
 
// This function can be marked `async` if using `await` inside
export function proxy(request: NextRequest) {
    console.log("middleware test");
    return NextResponse.next()
}

// 静的ファイルやAPIルートのリクエスト時には実行されないように設定
export const config = {
    matcher: [
      '/((?!api|_next/static|_next/image|favicon.ico|.*\\.(?:svg|png|jpg|jpeg|gif|webp|ico)).*)',
    ],
  };

// 特定のパスのみで実行する場合はこう設定
// export const config = {
//     matcher: '/about/:path*',
// }