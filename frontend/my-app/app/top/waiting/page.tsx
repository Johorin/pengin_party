'use client'

import { useState, useEffect } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import Link from 'next/link';
import { useCurrentUser } from '@/app/hooks/useCurrentUser';

const WaitingRoom = () => {
    const [roomId, setRoomId] = useState<string>('');
    const [participants, setParticipants] = useState<string[]>([]); // ホストを除く参加者リスト
    const [isHost, setIsHost] = useState<boolean>(false); // ホストかどうか
    const { user, loading } = useCurrentUser();
    const router = useRouter();
    const searchParams = useSearchParams();

    useEffect(() => {
        if (loading == true || !user) return;

        // URLパラメータからroomIdとisHostを取得
        const urlRoomId = searchParams.get('roomId');
        const urlIsHost = searchParams.get('isHost') === 'true';

        if (urlRoomId) {
            // 部屋に参加した場合
            setRoomId(urlRoomId);
            setIsHost(urlIsHost);
        } else {
            // 部屋を作成した場合
            async function createRoom(roomId: string) {
                const idToken = await user?.getIdToken();
                console.log("idToken: ", idToken);

                const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_API_ENDPOINT + '/rooms', {
                    method: 'POST',
                    headers: {
                        Authorization: `Bearer ${idToken}`,
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        "room_id": roomId
                    }),
                });

                const data = await res.json()
                console.log("fetch POST /rooms: ", data);
            }

            // 6桁の英数字IDを生成
            const generateRoomId = (): string => {
                const chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ';
                let result = '';
                for (let i = 0; i < 6; i++) {
                    result += chars.charAt(Math.floor(Math.random() * chars.length));
                }
                return result;
            };
            const newRoomId = generateRoomId();

            createRoom(newRoomId);
            setRoomId(newRoomId);
            setIsHost(true); // 部屋を作成した場合はホスト
            
            // URLを更新してroomIdとisHostを反映
            router.replace(`/top/waiting?roomId=${newRoomId}&isHost=true`);
        }
        
        // TODO: WebSocket等で参加者の更新を監視
        // 現時点ではUIのみの実装
    }, [user, loading, searchParams, router]);

    const handleStartGame = () => {
        // TODO: ゲーム開始処理を実装
        console.log('ゲームを開始:', roomId);
    };

    // ホスト名を取得
    const hostName = user?.displayName || user?.email || 'ホスト';

    // 最大3人まで表示（ホストを除く）
    const maxParticipants = 3;
    const participantSlots = Array.from({ length: maxParticipants }, (_, index) => 
        participants[index] || null
    );

    // 参加者が1人以上いるかチェック
    const hasParticipants = participants.length >= 1;

    return (
        <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 via-white to-purple-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 p-4">
            <div className="w-full max-w-md space-y-8">
                <div className="bg-white dark:bg-gray-800 rounded-2xl shadow-xl p-8 space-y-8">
                    {/* 部屋番号表示 */}
                    <div className="text-center space-y-2">
                        <p className="text-sm text-gray-600 dark:text-gray-400 mb-2">部屋番号</p>
                        {roomId && (
                            <p className="text-4xl font-bold tracking-widest text-gray-900 dark:text-white">
                                <span className="text-blue-600 dark:text-blue-400">{roomId}</span>
                            </p>
                        )}
                    </div>

                    {/* 参加者枠 */}
                    <div className="space-y-4">
                        <p className="text-sm font-medium text-gray-700 dark:text-gray-300 text-center">
                            参加者を待っています...
                        </p>
                        <div className="space-y-3">
                            {/* ホスト名の枠（一番上） */}
                            <div className="w-full h-20 rounded-xl border-2 border-blue-500 bg-blue-50 dark:bg-blue-900/20 dark:border-blue-400 flex items-center justify-center">
                                <div className="flex items-center space-x-3">
                                    <div className="w-10 h-10 rounded-full bg-gradient-to-br from-blue-400 to-purple-500 flex items-center justify-center">
                                        <span className="text-white font-bold text-lg">
                                            {hostName.charAt(0).toUpperCase()}
                                        </span>
                                    </div>
                                    <div className="flex flex-col">
                                        <span className="text-gray-900 dark:text-white font-medium">
                                            {hostName}
                                        </span>
                                        <span className="text-xs text-blue-600 dark:text-blue-400 font-semibold">
                                            ホスト
                                        </span>
                                    </div>
                                </div>
                            </div>

                            {/* 参加者の枠（最大3人） */}
                            {participantSlots.map((participant, index) => (
                                <div
                                    key={index}
                                    className={`w-full h-20 rounded-xl border-2 transition-all duration-200 ${
                                        participant
                                            ? 'border-green-500 bg-green-50 dark:bg-green-900/20 dark:border-green-400'
                                            : 'border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700/50'
                                    } flex items-center justify-center`}
                                >
                                    {participant ? (
                                        <div className="flex items-center space-x-3">
                                            <div className="w-10 h-10 rounded-full bg-gradient-to-br from-green-400 to-teal-500 flex items-center justify-center">
                                                <span className="text-white font-bold text-lg">
                                                    {participant.charAt(0).toUpperCase()}
                                                </span>
                                            </div>
                                            <span className="text-gray-900 dark:text-white font-medium">
                                                {participant}
                                            </span>
                                        </div>
                                    ) : (
                                        <div className="text-gray-400 dark:text-gray-500 text-sm">
                                            待機中...
                                        </div>
                                    )}
                                </div>
                            ))}
                        </div>
                    </div>

                    {/* ゲームを始めるボタン（ホストかつ参加者が1人以上いる時のみ表示） */}
                    {isHost && hasParticipants && (
                        <button
                            onClick={handleStartGame}
                            className="w-full bg-gradient-to-r from-purple-500 to-pink-600 hover:from-purple-600 hover:to-pink-700 text-white font-bold py-4 px-8 rounded-xl shadow-lg hover:shadow-xl transform hover:scale-105 transition-all duration-200 flex items-center justify-center space-x-3 group"
                        >
                            <svg className="w-5 h-5 group-hover:rotate-12 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z" />
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                            </svg>
                            <span className="text-lg">ゲームを始める</span>
                        </button>
                    )}

                    {/* 戻るボタン */}
                    <div className="text-center">
                        <Link href="/top">
                            <button className="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white transition-colors duration-200 flex items-center justify-center space-x-2 mx-auto">
                                <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 19l-7-7m0 0l7-7m-7 7h18" />
                                </svg>
                                <span>戻る</span>
                            </button>
                        </Link>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default WaitingRoom;
