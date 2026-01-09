'use client'

import { useCurrentUser } from "@/app/hooks/useCurrentUser";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
import Loading from "./Loading";

const AuthGuard = (props: { children: React.ReactNode }) => {
  const { user, loading } = useCurrentUser();
  const router = useRouter();
  
  console.log("AuthGuardです：", user, "loading:", loading);

  useEffect(() => {
    // ローディング中は何もしない
    if (loading) {
      return;
    }

    // 認証状態が確認された後、未ログインの場合はログインページへリダイレクト
    if (user == null) {
      router.push('/login');
    }
  }, [user, loading, router]);

  // ローディング中はローディング表示
  if (loading) {
    return <Loading />;
  }

  // 未ログインの場合は何も表示しない（リダイレクト処理中）
  if (user == null) {
    return null;
  }

  // ログイン中の場合のみ子のコンポーネントを描画
  return <>{props.children}</>
}

export default AuthGuard;