package mediashare

import (
	"html/template"
	"time"
)

// PageData contains data for rendering the player page.
type PageData struct {
	ID          string
	Filename    string
	ContentType string
	Size        int64
	IsVideo     bool
	IsAudio     bool
	IsImage     bool
	RawURL      string
	Username    string
	UploadedAt  time.Time
	ServiceName string
	Lang        string
	T           map[string]string
	BaseURL     string
}

// NotFoundData contains data for the 404 page.
type NotFoundData struct {
	ServiceName string
	Lang        string
	T           map[string]string
}

// FileListItem represents a file in the list view.
type FileListItem struct {
	ID           string
	Filename     string
	ContentType  string
	Username     string
	UploadedAt   string
	LastOpenedAt string
	URL          string
	RawURL       string
}

// ListPageData contains data for the file list page.
type ListPageData struct {
	ServiceName string
	Lang        string
	T           map[string]string
	Files       []FileListItem
	BaseURL     string
}

// MainImageData contains data for the main image page.
type MainImageData struct {
	ServiceName string
	Lang        string
	T           map[string]string
	ImagePath   string
}

const playerTemplateHTML = `<!DOCTYPE html>
<html lang="{{.Lang}}">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Filename}} - {{.ServiceName}}</title>

    <!-- Open Graph / Social Media -->
    <meta property="og:title" content="{{.Filename}}">
    <meta property="og:description" content="{{index .T "uploaded_by"}} {{.Username}}">
    <meta property="og:site_name" content="{{.ServiceName}}">
    {{if .IsVideo}}
    <meta property="og:type" content="video.other">
    <meta property="og:video" content="{{.RawURL}}">
    <meta property="og:video:type" content="{{.ContentType}}">
    {{else if .IsImage}}
    <meta property="og:type" content="website">
    <meta property="og:image" content="{{.RawURL}}">
    {{else if .IsAudio}}
    <meta property="og:type" content="music.song">
    <meta property="og:audio" content="{{.RawURL}}">
    {{end}}

    <!-- Twitter Card -->
    <meta name="twitter:card" content="{{if .IsVideo}}player{{else if .IsImage}}summary_large_image{{else}}summary{{end}}">

    <!-- Tailwind CSS -->
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        tailwind.config = {
            darkMode: 'class',
            theme: {
                extend: {
                    colors: {
                        primary: {
                            50: '#f0f9ff',
                            100: '#e0f2fe',
                            200: '#bae6fd',
                            300: '#7dd3fc',
                            400: '#38bdf8',
                            500: '#0ea5e9',
                            600: '#0284c7',
                            700: '#0369a1',
                            800: '#075985',
                            900: '#0c4a6e',
                        }
                    }
                }
            }
        }
    </script>

    <!-- Dark mode detection -->
    <script>
        if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
            document.documentElement.classList.add('dark');
        } else {
            document.documentElement.classList.remove('dark');
        }
    </script>

    <style>
        /* Custom scrollbar */
        ::-webkit-scrollbar { width: 8px; height: 8px; }
        ::-webkit-scrollbar-track { background: transparent; }
        ::-webkit-scrollbar-thumb { background: #888; border-radius: 4px; }
        ::-webkit-scrollbar-thumb:hover { background: #666; }

        /* Video/Audio player styling */
        video, audio {
            max-height: 70vh;
            border-radius: 0.75rem;
        }

        /* Image styling */
        .media-image {
            max-height: 80vh;
            object-fit: contain;
        }

        /* Animations */
        .fade-in {
            animation: fadeIn 0.3s ease-in-out;
        }
        @keyframes fadeIn {
            from { opacity: 0; transform: translateY(10px); }
            to { opacity: 1; transform: translateY(0); }
        }

        /* Copy button feedback */
        .copy-feedback {
            animation: copyPulse 0.5s ease-in-out;
        }
        @keyframes copyPulse {
            0%, 100% { transform: scale(1); }
            50% { transform: scale(1.1); }
        }
    </style>
</head>
<body class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-800 text-gray-900 dark:text-gray-100 transition-colors duration-300">

    <!-- Header -->
    <header class="fixed top-0 left-0 right-0 z-50 backdrop-blur-md bg-white/80 dark:bg-gray-900/80 border-b border-gray-200 dark:border-gray-700">
        <div class="max-w-6xl mx-auto px-4 py-3 flex items-center justify-between">
            <a href="/" class="flex items-center gap-2 text-lg font-semibold text-primary-600 dark:text-primary-400">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4v16M17 4v16M3 8h4m10 0h4M3 12h18M3 16h4m10 0h4M4 20h16a1 1 0 001-1V5a1 1 0 00-1-1H4a1 1 0 00-1 1v14a1 1 0 001 1z"/>
                </svg>
                {{.ServiceName}}
            </a>

            <!-- Theme toggle -->
            <button onclick="toggleTheme()" class="p-2 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors" title="Toggle theme">
                <svg class="w-5 h-5 hidden dark:block" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z" clip-rule="evenodd"/>
                </svg>
                <svg class="w-5 h-5 block dark:hidden" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"/>
                </svg>
            </button>
        </div>
    </header>

    <!-- Main content -->
    <main class="pt-20 pb-16 px-4">
        <div class="max-w-4xl mx-auto fade-in">

            <!-- Media container -->
            <div class="bg-white dark:bg-gray-800 rounded-2xl shadow-xl overflow-hidden">

                <!-- Player area -->
                <div class="relative bg-black flex items-center justify-center min-h-[300px]">
                    {{if .IsVideo}}
                    <video controls autoplay playsinline class="w-full" preload="metadata">
                        <source src="{{.RawURL}}" type="{{.ContentType}}">
                        {{index .T "unsupported"}}
                    </video>
                    {{else if .IsAudio}}
                    <div class="w-full p-8 flex flex-col items-center gap-6">
                        <div class="w-32 h-32 rounded-full bg-gradient-to-br from-primary-400 to-primary-600 flex items-center justify-center">
                            <svg class="w-16 h-16 text-white" fill="currentColor" viewBox="0 0 20 20">
                                <path d="M18 3a1 1 0 00-1.196-.98l-10 2A1 1 0 006 5v9.114A4.369 4.369 0 005 14c-1.657 0-3 .895-3 2s1.343 2 3 2 3-.895 3-2V7.82l8-1.6v5.894A4.37 4.37 0 0015 12c-1.657 0-3 .895-3 2s1.343 2 3 2 3-.895 3-2V3z"/>
                            </svg>
                        </div>
                        <audio controls autoplay class="w-full max-w-md" preload="metadata">
                            <source src="{{.RawURL}}" type="{{.ContentType}}">
                            {{index .T "unsupported"}}
                        </audio>
                    </div>
                    {{else if .IsImage}}
                    <a href="{{.RawURL}}" target="_blank" class="block">
                        <img src="{{.RawURL}}" alt="{{.Filename}}" class="media-image w-full" loading="eager">
                    </a>
                    {{else}}
                    <div class="p-8 text-center text-gray-400">
                        <svg class="w-16 h-16 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                        </svg>
                        <p>{{index .T "unsupported"}}</p>
                    </div>
                    {{end}}
                </div>

                <!-- Info section -->
                <div class="p-6 space-y-4">

                    <!-- Filename -->
                    <h1 class="text-xl font-semibold truncate" title="{{.Filename}}">{{.Filename}}</h1>

                    <!-- Meta info -->
                    <div class="flex flex-wrap gap-4 text-sm text-gray-600 dark:text-gray-400">
                        {{if .Username}}
                        <div class="flex items-center gap-2">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                            </svg>
                            <span>{{index .T "uploaded_by"}} <strong class="text-gray-900 dark:text-gray-100">{{.Username}}</strong></span>
                        </div>
                        {{end}}
                        <div class="flex items-center gap-2">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
                            </svg>
                            <span>{{index .T "uploaded_at"}} <strong class="text-gray-900 dark:text-gray-100">{{.UploadedAt.Format "2006-01-02 15:04"}}</strong></span>
                        </div>
                    </div>

                    <!-- Actions -->
                    <div class="flex flex-wrap gap-3 pt-2">
                        <a href="{{.RawURL}}" download="{{.Filename}}"
                           class="inline-flex items-center gap-2 px-4 py-2 bg-primary-600 hover:bg-primary-700 text-white font-medium rounded-lg transition-colors">
                            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/>
                            </svg>
                            {{index .T "download"}}
                        </a>

                        <button onclick="copyLink()" id="copyBtn"
                                class="inline-flex items-center gap-2 px-4 py-2 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 font-medium rounded-lg transition-colors">
                            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3"/>
                            </svg>
                            <span id="copyText">{{index .T "copy_link"}}</span>
                        </button>

                        <a href="{{.RawURL}}" target="_blank"
                           class="inline-flex items-center gap-2 px-4 py-2 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 font-medium rounded-lg transition-colors">
                            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"/>
                            </svg>
                            {{index .T "open_in_new_tab"}}
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </main>

    <!-- Footer -->
    <footer class="fixed bottom-0 left-0 right-0 py-3 text-center text-sm text-gray-500 dark:text-gray-400 bg-white/80 dark:bg-gray-900/80 backdrop-blur-md border-t border-gray-200 dark:border-gray-700">
        {{index .T "powered_by"}} <strong>{{.ServiceName}}</strong>
    </footer>

    <script>
        function toggleTheme() {
            document.documentElement.classList.toggle('dark');
            localStorage.theme = document.documentElement.classList.contains('dark') ? 'dark' : 'light';
        }

        function copyLink() {
            navigator.clipboard.writeText(window.location.href).then(() => {
                const btn = document.getElementById('copyBtn');
                const text = document.getElementById('copyText');
                const originalText = text.textContent;
                text.textContent = '{{index .T "copied"}}';
                btn.classList.add('copy-feedback');
                setTimeout(() => {
                    text.textContent = originalText;
                    btn.classList.remove('copy-feedback');
                }, 2000);
            });
        }
    </script>
</body>
</html>`

const notFoundTemplateHTML = `<!DOCTYPE html>
<html lang="{{.Lang}}">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{index .T "not_found"}} - {{.ServiceName}}</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        tailwind.config = {
            darkMode: 'class',
            theme: {
                extend: {
                    colors: {
                        primary: {
                            600: '#0284c7',
                            700: '#0369a1',
                        }
                    }
                }
            }
        }
        if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
            document.documentElement.classList.add('dark');
        }
    </script>
</head>
<body class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-800 text-gray-900 dark:text-gray-100 flex items-center justify-center p-4">
    <div class="text-center">
        <div class="mb-8">
            <svg class="w-24 h-24 mx-auto text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
        </div>
        <h1 class="text-3xl font-bold mb-2">{{index .T "not_found"}}</h1>
        <p class="text-gray-600 dark:text-gray-400 mb-8">{{index .T "not_found_desc"}}</p>
        <a href="/" class="inline-flex items-center gap-2 px-6 py-3 bg-primary-600 hover:bg-primary-700 text-white font-medium rounded-lg transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
            </svg>
            {{.ServiceName}}
        </a>
    </div>
</body>
</html>`

const listTemplateHTML = `<!DOCTYPE html>
<html lang="{{.Lang}}">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{index .T "recent_files"}} - {{.ServiceName}}</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        tailwind.config = {
            darkMode: 'class',
            theme: {
                extend: {
                    colors: {
                        primary: {
                            50: '#f0f9ff',
                            100: '#e0f2fe',
                            500: '#0ea5e9',
                            600: '#0284c7',
                            700: '#0369a1',
                        }
                    }
                }
            }
        }
        if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
            document.documentElement.classList.add('dark');
        }
    </script>
    <style>
        .fade-in { animation: fadeIn 0.3s ease-in-out; }
        @keyframes fadeIn {
            from { opacity: 0; transform: translateY(10px); }
            to { opacity: 1; transform: translateY(0); }
        }
        /* Mini player styles */
        .mini-video { max-height: 120px; border-radius: 0.5rem; }
        .mini-audio { width: 100%; max-width: 250px; }
        .mini-thumb { max-height: 80px; max-width: 120px; object-fit: cover; border-radius: 0.5rem; cursor: pointer; transition: transform 0.2s; }
        .mini-thumb:hover { transform: scale(1.05); }
        /* Lightbox */
        .lightbox { display: none; position: fixed; inset: 0; background: rgba(0,0,0,0.9); z-index: 100; justify-content: center; align-items: center; }
        .lightbox.active { display: flex; }
        .lightbox img { max-width: 90vw; max-height: 90vh; object-fit: contain; }
        .lightbox-close { position: absolute; top: 1rem; right: 1rem; color: white; font-size: 2rem; cursor: pointer; }
    </style>
</head>
<body class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-800 text-gray-900 dark:text-gray-100 transition-colors duration-300">

    <!-- Header -->
    <header class="fixed top-0 left-0 right-0 z-50 backdrop-blur-md bg-white/80 dark:bg-gray-900/80 border-b border-gray-200 dark:border-gray-700">
        <div class="max-w-6xl mx-auto px-4 py-3 flex items-center justify-between">
            <a href="/" class="flex items-center gap-2 text-lg font-semibold text-primary-600 dark:text-primary-400">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4v16M17 4v16M3 8h4m10 0h4M3 12h18M3 16h4m10 0h4M4 20h16a1 1 0 001-1V5a1 1 0 00-1-1H4a1 1 0 00-1 1v14a1 1 0 001 1z"/>
                </svg>
                {{.ServiceName}}
            </a>
            <button onclick="toggleTheme()" class="p-2 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors" title="Toggle theme">
                <svg class="w-5 h-5 hidden dark:block" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z" clip-rule="evenodd"/>
                </svg>
                <svg class="w-5 h-5 block dark:hidden" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"/>
                </svg>
            </button>
        </div>
    </header>

    <!-- Lightbox for images -->
    <div id="lightbox" class="lightbox" onclick="closeLightbox()">
        <span class="lightbox-close">&times;</span>
        <img id="lightbox-img" src="" alt="">
    </div>

    <!-- Main content -->
    <main class="pt-20 pb-16 px-4">
        <div class="max-w-6xl mx-auto fade-in">
            <h1 class="text-2xl font-bold mb-6">{{index .T "recent_files"}}</h1>

            {{if .Files}}
            <div class="grid gap-4">
                {{range .Files}}
                <div class="bg-white dark:bg-gray-800 rounded-xl shadow-lg overflow-hidden hover:shadow-xl transition-shadow">
                    <div class="flex flex-col md:flex-row">
                        <!-- Media preview -->
                        <div class="md:w-64 bg-gray-100 dark:bg-gray-900 flex items-center justify-center p-3 min-h-[100px]">
                            {{if eq .ContentType "video"}}
                            <video controls preload="metadata" class="mini-video w-full" poster="">
                                <source src="{{.RawURL}}#t=0.5" type="video/mp4">
                            </video>
                            {{else if eq .ContentType "audio"}}
                            <div class="w-full flex flex-col items-center gap-2">
                                <svg class="w-8 h-8 text-purple-500" fill="currentColor" viewBox="0 0 20 20"><path d="M18 3a1 1 0 00-1.196-.98l-10 2A1 1 0 006 5v9.114A4.369 4.369 0 005 14c-1.657 0-3 .895-3 2s1.343 2 3 2 3-.895 3-2V7.82l8-1.6v5.894A4.37 4.37 0 0015 12c-1.657 0-3 .895-3 2s1.343 2 3 2 3-.895 3-2V3z"/></svg>
                                <audio controls preload="metadata" class="mini-audio">
                                    <source src="{{.RawURL}}">
                                </audio>
                            </div>
                            {{else if eq .ContentType "image"}}
                            <img src="{{.RawURL}}" alt="{{.Filename}}" class="mini-thumb" onclick="openLightbox('{{.RawURL}}', event)" loading="lazy">
                            {{else}}
                            <div class="text-center text-gray-400">
                                <svg class="w-10 h-10 mx-auto" fill="currentColor" viewBox="0 0 20 20"><path fill-rule="evenodd" d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z" clip-rule="evenodd"/></svg>
                            </div>
                            {{end}}
                        </div>

                        <!-- File info -->
                        <div class="flex-1 p-4">
                            <div class="flex flex-wrap items-start justify-between gap-2">
                                <div class="flex-1 min-w-0">
                                    <a href="{{.URL}}" class="text-lg font-semibold text-gray-900 dark:text-gray-100 hover:text-primary-600 dark:hover:text-primary-400 truncate block" title="{{.Filename}}">
                                        {{.Filename}}
                                    </a>
                                    <div class="flex flex-wrap items-center gap-x-4 gap-y-1 mt-2 text-sm text-gray-600 dark:text-gray-400">
                                        <span class="flex items-center gap-1">
                                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                                            {{.Username}}
                                        </span>
                                        <span class="flex items-center gap-1">
                                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
                                            {{.UploadedAt}}
                                        </span>
                                        <span class="flex items-center gap-1 text-xs">
                                            <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/></svg>
                                            {{.LastOpenedAt}}
                                        </span>
                                    </div>
                                </div>
                                <div class="flex items-center gap-2">
                                    {{if eq .ContentType "video"}}
                                    <span class="px-2 py-1 text-xs font-medium bg-blue-100 text-blue-700 dark:bg-blue-900 dark:text-blue-300 rounded">VIDEO</span>
                                    {{else if eq .ContentType "audio"}}
                                    <span class="px-2 py-1 text-xs font-medium bg-purple-100 text-purple-700 dark:bg-purple-900 dark:text-purple-300 rounded">AUDIO</span>
                                    {{else if eq .ContentType "image"}}
                                    <span class="px-2 py-1 text-xs font-medium bg-green-100 text-green-700 dark:bg-green-900 dark:text-green-300 rounded">IMAGE</span>
                                    {{else}}
                                    <span class="px-2 py-1 text-xs font-medium bg-gray-100 text-gray-700 dark:bg-gray-700 dark:text-gray-300 rounded">FILE</span>
                                    {{end}}
                                    <a href="{{.URL}}" class="p-2 text-primary-600 dark:text-primary-400 hover:bg-primary-50 dark:hover:bg-primary-900/20 rounded-lg transition-colors" title="{{index $.T "open_in_new_tab"}}">
                                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"/></svg>
                                    </a>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                {{end}}
            </div>
            {{else}}
            <div class="bg-white dark:bg-gray-800 rounded-2xl shadow-xl p-12 text-center">
                <svg class="w-16 h-16 mx-auto text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"/>
                </svg>
                <h2 class="text-xl font-semibold mb-2">{{index .T "no_files"}}</h2>
                <p class="text-gray-600 dark:text-gray-400">{{index .T "no_files_desc"}}</p>
            </div>
            {{end}}
        </div>
    </main>

    <!-- Footer -->
    <footer class="fixed bottom-0 left-0 right-0 py-3 text-center text-sm text-gray-500 dark:text-gray-400 bg-white/80 dark:bg-gray-900/80 backdrop-blur-md border-t border-gray-200 dark:border-gray-700">
        {{index .T "powered_by"}} <strong>{{.ServiceName}}</strong>
    </footer>

    <script>
        function toggleTheme() {
            document.documentElement.classList.toggle('dark');
            localStorage.theme = document.documentElement.classList.contains('dark') ? 'dark' : 'light';
        }

        function openLightbox(src, event) {
            event.stopPropagation();
            document.getElementById('lightbox-img').src = src;
            document.getElementById('lightbox').classList.add('active');
        }

        function closeLightbox() {
            document.getElementById('lightbox').classList.remove('active');
        }

        document.addEventListener('keydown', function(e) {
            if (e.key === 'Escape') closeLightbox();
        });
    </script>
</body>
</html>`

const mainImageTemplateHTML = `<!DOCTYPE html>
<html lang="{{.Lang}}">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.ServiceName}}</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        tailwind.config = {
            darkMode: 'class',
            theme: {
                extend: {
                    colors: {
                        primary: {
                            600: '#0284c7',
                            700: '#0369a1',
                        }
                    }
                }
            }
        }
        if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
            document.documentElement.classList.add('dark');
        }
    </script>
    <style>
        .fade-in { animation: fadeIn 0.5s ease-in-out; }
        @keyframes fadeIn {
            from { opacity: 0; transform: scale(0.95); }
            to { opacity: 1; transform: scale(1); }
        }
    </style>
</head>
<body class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-800 flex items-center justify-center p-4">
    <div class="text-center fade-in">
        <div class="bg-white dark:bg-gray-800 rounded-2xl shadow-2xl p-6 max-w-2xl mx-auto">
            <img src="{{.ImagePath}}" alt="{{.ServiceName}}" class="max-w-full max-h-[70vh] mx-auto rounded-xl shadow-lg object-contain">
            <h1 class="mt-6 text-2xl font-bold text-gray-900 dark:text-gray-100">{{.ServiceName}}</h1>
        </div>
    </div>

    <button onclick="toggleTheme()" class="fixed top-4 right-4 p-2 rounded-lg bg-white/80 dark:bg-gray-800/80 hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors shadow-lg" title="Toggle theme">
        <svg class="w-5 h-5 hidden dark:block text-gray-100" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z" clip-rule="evenodd"/>
        </svg>
        <svg class="w-5 h-5 block dark:hidden text-gray-900" fill="currentColor" viewBox="0 0 20 20">
            <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"/>
        </svg>
    </button>

    <script>
        function toggleTheme() {
            document.documentElement.classList.toggle('dark');
            localStorage.theme = document.documentElement.classList.contains('dark') ? 'dark' : 'light';
        }
    </script>
</body>
</html>`

var (
	// PlayerTemplate is the parsed template for the player page.
	PlayerTemplate *template.Template

	// NotFoundTemplate is the parsed template for 404 errors.
	NotFoundTemplate *template.Template

	// ListTemplate is the parsed template for the file list page.
	ListTemplate *template.Template

	// MainImageTemplate is the parsed template for the main image page.
	MainImageTemplate *template.Template
)

func init() {
	PlayerTemplate = template.Must(template.New("player").Parse(playerTemplateHTML))
	NotFoundTemplate = template.Must(template.New("notfound").Parse(notFoundTemplateHTML))
	ListTemplate = template.Must(template.New("list").Parse(listTemplateHTML))
	MainImageTemplate = template.Must(template.New("mainimage").Parse(mainImageTemplateHTML))
}
