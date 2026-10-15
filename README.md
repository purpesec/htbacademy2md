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
```

### Idiomas soportados (`-lang`)

HTB Academy ofrece traducción automática (IA) del contenido a: `hi`, `ar`, `es`, `fr`, `zh` (simplificado, HK y TW), `el`, `ja`.
Sin `-lang`, se descarga el contenido original en inglés.

### Obtener la cookie

1. Inicia sesión en [HTB Academy](https://academy.hackthebox.com) en tu navegador.
2. Abre las DevTools (F12) → Application/Storage → Cookies → `academy.hackthebox.com`.
3. Copia el valor de la cookie `htb_academy_session`.

Puedes pasarla con `-c "htb_academy_session=VALOR"`; también puedes incluir otras cookies separadas por `;`.

## Notas

- Pensada para uso personal: descarga únicamente módulos a los que tengas acceso y no compartas el contenido sin autorización de HackTheBox.
- Los títulos de módulo y secciones que devuelve la API de HTB permanecen en inglés; el contenido de las secciones sí llega traducido.

## Créditos

Basado en [htb-academy-to-md](https://github.com/Tut-k0/htb-academy-to-md).

## Licencia

[MIT](LICENSE)
