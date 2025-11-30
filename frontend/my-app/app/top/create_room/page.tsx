'use client'

import { useState, useEffect } from 'react';

const CreateRoom = () => {
    const [roomUid, setRoomUid] = useState<string>('');

    useEffect(() => {
        // 6桁の英数字UIDを生成
        const generateRoomUid = (): string => {
            const chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ';
            let uid = '';
            for (let i = 0; i < 6; i++) {
                uid += chars.charAt(Math.floor(Math.random() * chars.length));
            }
            return uid;
        };

        setRoomUid(generateRoomUid());
    }, []);

    return (
        <div className="random-room-uid">
            {roomUid && (
                <div className="text-center">
                    <p className="text-2xl font-bold text-gray-900 dark:text-white">
                        ルームID: <span className="text-blue-600 dark:text-blue-400">{roomUid}</span>
                    </p>
                </div>
            )}
        </div>
    );
};

export default CreateRoom;