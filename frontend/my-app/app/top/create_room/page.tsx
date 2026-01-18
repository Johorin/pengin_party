'use client'

import { useState, useEffect } from 'react';
import { useCurrentUser } from '@/app/hooks/useCurrentUser';

const CreateRoom = () => {
    const [roomId, setRoomId] = useState<string>('');
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
    }, [user, loading]);

    return (
        <div className="random-room-uid">
            {roomId && (
                <div className="text-center">
                    <p className="text-2xl font-bold text-gray-900 dark:text-white">
                        ルームID: <span className="text-blue-600 dark:text-blue-400">{roomId}</span>
                    </p>
                </div>
            )}
        </div>
    );
};

export default CreateRoom;