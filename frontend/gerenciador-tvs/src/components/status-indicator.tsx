interface StatusIndicatorProps {
  isOnline: boolean
  isChecking: boolean
}

export function StatusIndicator({ isOnline, isChecking }: StatusIndicatorProps) {
  return (
    <div className="flex items-center gap-2">
      <div
        className={`h-3 w-3 rounded-full ${
          isChecking
            ? 'bg-yellow-500 animate-pulse'
            : isOnline
            ? 'bg-green-500'
            : 'bg-red-500'
        }`}
      />
      <span className="text-sm text-white">
        {isChecking ? 'Verificando...' : isOnline ? 'Online' : 'Offline'}
      </span>
    </div>
  )
}

