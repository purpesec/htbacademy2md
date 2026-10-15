# htbacademy2md

Herramienta CLI en Go para exportar módulos de **HTB Academy** a Markdown, ideal para estudiar sin conexión o tomar notas en Obsidian u otro editor.

## Características

- Exporta un módulo completo (todas sus secciones) a un único archivo `.md`.
- **Soporte multi-idioma**: descarga el contenido traducido por la IA de HTB Academy (`-lang es`, `fr`, `ja`, ...).
- Opción de descargar las imágenes localmente (`-local_images`) o dejarlas referenciadas por URL.
- Autenticación mediante cookie de sesión (`htb_academy_session`), sin necesidad de contraseña ni 2FA.

## Compilar

Requiere Go 1.25 o superior.

```bash
cd src
go build -o htbacademy2md
```

## Uso

```bash
# Mostrar la ayuda
htbacademy2md -h

# Exportar un módulo (contenido original en inglés, imágenes por URL)
htbacademy2md -m https://academy.hackthebox.com/module/112/section/1060 -c "htb_academy_session=..."

# Exportar en español con imágenes locales
htbacademy2md -m https://academy.hackthebox.com/module/112/section/1060 -lang es -local_images -c "htb_academy_session=..."

# Si ya tienes un .env configurado (ver más abajo), basta con la URL
htbacademy2md -m https://academy.hackthebox.com/module/112/section/1060
```

### Idiomas soportados (`-lang`)

HTB Academy ofrece traducción automática (IA) del contenido a: `hi`, `ar`, `es`, `fr`, `zh` (simplificado, HK y TW), `el`, `ja`.
Sin `-lang`, se descarga el contenido original en inglés.

### Obtener la cookie

1. Inicia sesión en [HTB Academy](https://academy.hackthebox.com) en tu navegador.
2. Abre las DevTools (F12) → Application/Storage → Cookies → `academy.hackthebox.com`.
3. Copia el valor de la cookie `htb_academy_session`.

Puedes pasarla con `-c "htb_academy_session=VALOR"`; también puedes incluir otras cookies separadas por `;`.

### Configuración con `.env` (opcional)

Guarda tus cookies y valores por defecto en un archivo `.env` para no pasarlos en cada comando. Usa `.env.example` como plantilla:

```bash
cp .env.example .env
# edita .env con tus valores
```

| Variable | Descripción | Equivale a |
|---|---|---|
| `HTB_COOKIES` | Cookies de autenticación | `-c` |
| `HTB_LANG` | Idioma por defecto del contenido | `-lang` |
| `HTB_LOCAL_IMAGES` | `true`/`false`: descargar imágenes localmente | `-local_images` |

Ubicaciones buscadas (se usa el primero que exista):

1. La ruta indicada con `-env RUTA`
2. `./.env` (carpeta actual)
3. `~/.config/htbacademy2md/.env` (configuración global del usuario)

Los flags de la línea de comandos siempre tienen prioridad sobre el `.env`.

> **Importante**: `.env` está en `.gitignore`, así que tus cookies nunca se subirán al repositorio.

## Notas

- Pensada para uso personal: descarga únicamente módulos a los que tengas acceso y no compartas el contenido sin autorización de HackTheBox.
- Los títulos de módulo y secciones que devuelve la API de HTB permanecen en inglés; el contenido de las secciones sí llega traducido.

## Créditos

Basado en [htb-academy-to-md](https://github.com/Tut-k0/htb-academy-to-md).

## Licencia

[MIT](LICENSE)
