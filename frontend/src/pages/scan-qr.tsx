/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable react-hooks/exhaustive-deps */

import React, { useState, useEffect, useRef } from 'react';
import styled from 'styled-components';
import { useRouter } from 'next/router';
import Image from 'next/image';
import Webcam from 'react-webcam';
import jsQR from 'jsqr';
import withAuth from '@/components/withAuth';

const ScanQRPage: React.FC = () => {
    const router = useRouter();
    const { eventId } = router.query;
    const webcamRef = useRef<Webcam>(null);
    const canvasRef = useRef<HTMLCanvasElement>(null);
    const [data, setData] = useState('No result');
    const [isActive, setIsActive] = useState(false);
    const [isLoading, setIsLoading] = useState(true);
    const [qrWorker, setQrWorker] = useState<Worker | null>(null);

    useEffect(() => {
        if (typeof window !== 'undefined') {
            const worker = new Worker(
                new URL('../workers/qr-worker.ts', import.meta.url)
            );
            setQrWorker(worker);

            return () => {
                worker.terminate();
            };
        }
    }, []);

    useEffect(() => {
        if (qrWorker) {
            qrWorker.onmessage = (event) => {
                if (event.data) {
                    handleScan(event.data);
                }
            };
        }
    }, [qrWorker]);

    useEffect(() => {
        setIsActive(true);
        const scanInterval = setInterval(scanQRCode, 100); // Scan every 100ms

        return () => {
            setIsActive(false);
            clearInterval(scanInterval);
            console.log('Cleanup called, QR Reader deactivated');
        };
    }, []);

    const scanQRCode = () => {
        if (webcamRef.current && canvasRef.current && qrWorker) {
            const video = webcamRef.current.video;
            const canvas = canvasRef.current;
            const context = canvas.getContext('2d', { willReadFrequently: true });

            if (video && video.readyState === 4 && context) {
                setIsLoading(false);
                const centerX = video.videoWidth / 2;
                const centerY = video.videoHeight / 2;
                const scanSize = Math.min(video.videoWidth, video.videoHeight) * 0.7;
                const startX = centerX - scanSize / 2;
                const startY = centerY - scanSize / 2;

                canvas.width = scanSize;
                canvas.height = scanSize;
                context.drawImage(video, startX, startY, scanSize, scanSize, 0, 0, scanSize, scanSize);
                const imageData = context.getImageData(0, 0, scanSize, scanSize);
                
                qrWorker.postMessage({
                    data: imageData.data,
                    width: imageData.width,
                    height: imageData.height
                });
            }
        }
    };

    const handleScan = (data: string) => {
        if (data) {
            setData(data);
            console.log('QR Code Data:', data);
            // Currently, no API call for verification
        }
    };

    const handleError = (error: any) => {
        console.error('QR Code Scan Error:', error);
        setIsLoading(false);
    };

    const handleGoBack = () => {
        setIsActive(false);
        router.push('/events');
    };

    return (
        <MainContainer>
            <Card>
                <Header>
                    <GoBackButton onClick={handleGoBack}>
                        <Image src="/left-arrow.svg" alt="go back" width={20} height={20} />
                        <ButtonTitle>Events</ButtonTitle>
                    </GoBackButton>
                </Header>
                <Title>Scan QR Code</Title>
                <ScannerContainer>
                    {isLoading && <LoadingText>Initializing camera...</LoadingText>}
                    {isActive && (
                        <>
                            <Webcam
                                ref={webcamRef}
                                audio={false}
                                screenshotFormat="image/jpeg"
                                videoConstraints={{ 
                                    facingMode: 'environment',
                                    width: { ideal: 640 },
                                    height: { ideal: 480 }
                                }}
                                onUserMedia={() => setIsLoading(false)}
                                onUserMediaError={handleError}
                                style={{ 
                                    width: '100%', 
                                    height: '300px', 
                                    objectFit: 'cover',
                                    borderRadius: '12px',
                                    display: isLoading ? 'none' : 'block'
                                }}
                            />
                            <Canvas ref={canvasRef} />
                            <ScanOverlay />
                        </>
                    )}
                </ScannerContainer>
                <Result>{data}</Result>
            </Card>
        </MainContainer>
    );
};

const MainContainer = styled.div`
    background-color: #121212;
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    padding: 20px;
    font-family: "Inter", sans-serif;
`;

const Card = styled.div`
    background-color: #1e1e1e;
    color: #fff;
    border-radius: 16px;
    padding: 24px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
    width: 100%;
    max-width: 480px;
`;

const Header = styled.header`
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 24px;
`;

const GoBackButton = styled.button`
    background: none;
    border: none;
    cursor: pointer;
    display: flex;
    align-items: center;
    color: #fff;
    font-size: 16px;
    padding: 0;
`;

const ButtonTitle = styled.span`
    font-size: 18px;
    font-weight: 600;
    background: linear-gradient(45deg, #ff6b6b, #feca57);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
`;

const Title = styled.h1`
    font-size: 24px;
    color: white;
    font-weight: 600;
    text-align: center;
    margin-bottom: 24px;
    background: linear-gradient(45deg, #ff6b6b, #feca57);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
`;

const ScannerContainer = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #1e1e1e;
    border-radius: 16px;
    margin-bottom: 24px;
    overflow: hidden;
    position: relative;
`;

const Canvas = styled.canvas`
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
`;

const Result = styled.p`
    color: rgba(255, 255, 255, 0.8);
    text-align: center;
    margin: 0;
    font-size: 16px;
    line-height: 1.6;
`;

const LoadingText = styled.p`
    color: white;
    text-align: center;
`;

const ScanOverlay = styled.div`
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 70%;
    height: 70%;
    border: 2px solid #fff;
    border-radius: 20px;
    box-shadow: 0 0 0 9999px rgba(0, 0, 0, 0.5);
`;

export default withAuth(ScanQRPage);
