'use client'

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { useCurrentUser } from '@/app/hooks/useCurrentUser';

const JoinRoom = () => {
    const [roomId, setRoomId] = useState<string>('');
    const [error, setError] = useState<string>('');
    const [loading, setLoading] = useState<boolean>(false);
    const router = useRouter();
    const { user } = useCurrentUser();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        
        if (!roomId.trim()) {
            setError('部屋IDを入力してください');
            return;
        }

        if (roomId.length !== 6) {
            setError('部屋IDは6桁の英数字です');
            return;
        }

        if (!user) {
            setError('ログインが必要です');
            return;
        }

        setLoading(true);
        setError('');

        try {
            const idToken = await user.getIdToken();
            
            const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_API_ENDPOINT + '/rooms/join', {
                method: 'POST',
                headers: {
                    Authorization: `Bearer ${idToken}`,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    room_id: roomId
                }),
            });

            if (!res.ok) {
                const errorData = await res.json();
                throw new Error(errorData.error || '部屋への参加に失敗しました');
            }

            const data = await res.json();
            console.log('部屋参加成功:', data);
            
            // 部屋参加成功後、待機画面に遷移（ホストではない）
            router.push(`/top/waiting?roomId=${roomId}&isHost=false`);
        } catch (err: any) {
            setError(err.message || '部屋への参加に失敗しました');
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 via-white to-purple-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 p-4">
            <div className="w-full max-w-md space-y-6">
                <div className="bg-white dark:bg-gray-800 rounded-2xl shadow-xl p-8 space-y-6">
                    <div className="text-center space-y-2">
                        <div className="inline-block p-4 bg-gradient-to-br from-green-500 to-teal-600 rounded-2xl mb-4">
                            <svg className="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1" />
                            </svg>
                        </div>
                        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
                            部屋に入る
                        </h1>
                        <p className="text-gray-600 dark:text-gray-400">
                            部屋IDを入力してください
                        </p>
                    </div>

                    <form onSubmit={handleSubmit} className="space-y-4">
                        <div>
                            <label htmlFor="roomId" className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                                部屋ID
                            </label>
                            <input
                                id="roomId"
                                type="text"
                                value={roomId}
                                onChange={(e) => {
                                    const value = e.target.value.toUpperCase().replace(/[^0-9A-Z]/g, '');
                                    setRoomId(value);
                                    setError('');
                                }}
                                maxLength={6}
                                placeholder="例: ABC123"
                                className="w-full px-4 py-3 text-center text-2xl font-bold tracking-widest bg-gray-50 dark:bg-gray-700 border-2 border-gray-300 dark:border-gray-600 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent text-gray-900 dark:text-white transition-all duration-200"
                            />
                        </div>

                        {error && (
                            <div className="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-4">
                                <p className="text-sm text-red-600 dark:text-red-400">{error}</p>
                            </div>
                        )}

                        <button
                            type="submit"
                            disabled={loading}
                            className="w-full bg-gradient-to-r from-green-500 to-teal-600 hover:from-green-600 hover:to-teal-700 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-4 px-8 rounded-xl shadow-lg hover:shadow-xl transform hover:scale-105 transition-all duration-200 flex items-center justify-center space-x-3 group"
                        >
                            {loading ? (
                                <>
                                    <div className="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></div>
                                    <span className="text-lg">参加中...</span>
                                </>
                            ) : (
                                <>
                                    <svg className="w-5 h-5 group-hover:translate-x-1 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1" />
                                    </svg>
                                    <span className="text-lg">参加する</span>
                                </>
                            )}
                        </button>
                    </form>

                    <div className="text-center">
                        <button
                            onClick={() => router.back()}
                            className="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white transition-colors duration-200"
                        >
                            戻る
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default JoinRoom;
