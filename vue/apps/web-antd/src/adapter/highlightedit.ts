import { ref, h, computed, defineComponent, watch } from 'vue'
import hljs from 'highlight.js/lib/core'

const component = defineComponent({
    props: {
        code: {
            type: String,
            required: true,
        },
        language: {
            type: String,
            default: '',
        },
        autodetect: {
            type: Boolean,
            default: true,
        },
        ignoreIllegals: {
            type: Boolean,
            default: true,
        },
    },
    setup(props, { emit }) {
        const language = ref(props.language)
        const ignoreIllegals = ref(props.ignoreIllegals)
        watch(() => props.language, (newLanguage) => {
            language.value = newLanguage
        })
        watch(() => props.ignoreIllegals, (newIgnoreIllegals) => {
            ignoreIllegals.value = newIgnoreIllegals
        })

        const autodetect = computed(() => props.autodetect || !language.value)
        const cannotDetectLanguage = computed(() => !autodetect.value && !hljs.getLanguage(language.value))

        const className = computed((): string => {
            if (cannotDetectLanguage.value) {
                return ''
            } else {
                return `hljs ${language.value}`
            }
        })

        const highlightedCode = computed((): string => {
            // No idea what language to use, return raw code
            if (cannotDetectLanguage.value) {
                console.warn(`The language "${language.value}" you specified could not be found.`)
                return (props.code)
            }

            if (autodetect.value) {
                const result = hljs.highlightAuto(props.code)
                language.value = result.language ?? ''
                return result.value
            } else {
                const result = hljs.highlight(props.code, {
                    language: language.value,
                    ignoreIllegals: props.ignoreIllegals,
                })
                return result.value
            }
        })
        const saveSelection = (e:any) => {
          const selection = window.getSelection();
          if (selection && selection.rangeCount && selection.rangeCount > 0) {
            const range = selection.getRangeAt(0)
            return range
          }
          return null;
        }
        const restoreSelection = (range:any, el:any) => {
          if (range) {
            const selection = window.getSelection();
            if (selection) {
              selection.removeAllRanges()
              selection.addRange(range)
              console.log('selection range', range)
            }
          }
        }
        const updateContent = (e:any) => {
          // const range = saveSelection(e)
          const result = hljs.highlight(e.target.innerText, {
            language: language.value,
            ignoreIllegals: props.ignoreIllegals,
          });
          e.target.innerHTML = result.value
          emit('update:code', e.target.innerText)
          // setTimeout(() => {
          //   restoreSelection(range, e.target)
          // }, 0)
        }
        return {
            className,
            highlightedCode,
            updateContent,
            emit,
        }
    },
    render() {
        const updateContent = this.updateContent;
        return h('pre', {}, [
            h('code', {
                class: this.className,
                innerHTML: this.highlightedCode,
                contenteditable:true,
                onInput(e: any) {
                },
                onBlur(e: any) {
                  updateContent(e);
                },
            }),
        ])
    },
})

export default component
