'use client'

import { useState, useEffect } from 'react';
import { useCurrentUser } from '@/app/hooks/useCurrentUser';

const CreateRoom = () => {
    const [roomId, setRoomId] = useState<string>('');
    const [participants, setParticipants] = useState<string[]>([]); // ホストを除く参加者リスト
    const { user, loading } = useCurrentUser();

    useEffect(() => {
        if (loading == true || !user) return;

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
        const roomId = generateRoomId();

        createRoom(roomId);
        setRoomId(roomId);
        
        // TODO: WebSocket等で参加者の更新を監視
        // 現時点ではUIのみの実装
    }, [user, loading]);

    // 最大3人まで表示（ホストを除く）
    const maxParticipants = 3;
    const participantSlots = Array.from({ length: maxParticipants }, (_, index) => 
        participants[index] || null
    );

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
                </div>
            </div>
        </div>
    );
};

export default CreateRoom;