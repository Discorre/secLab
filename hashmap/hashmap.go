package hashmap

// Node представляет элемент связного списка.
type Node struct {
	key   rune
	value rune
	next  *Node
}

// HashMap представляет хэш-таблицу.
type HashMap struct {
	buckets []*Node
	size    int
}

// NewHashMap создает новую хэш-таблицу с заданным размером.
func NewHashMap() *HashMap {
	return &HashMap{
		buckets: make([]*Node, 100), // Устанавливаем размер по умолчанию
		size:    100,
	}
}

// hash функция для вычисления индекса на основе ключа.
func (h *HashMap) hash(key rune) int {
	return int(key) % h.size
}

// Put добавляет значение в HashMap.
func (h *HashMap) Put(key rune, value rune) {
	index := h.hash(key)
	newNode := &Node{key: key, value: value}

	if h.buckets[index] == nil {
		h.buckets[index] = newNode
	} else {
		current := h.buckets[index]
		for current != nil {
			if current.key == key {
				current.value = value // Обновляем значение, если ключ уже существует
				return
			}
			if current.next == nil {
				break
			}
			current = current.next
		}
		current.next = newNode
	}
}

// Get получает значение по ключу.
func (h *HashMap) Get(key rune) (rune, bool) {
	index := h.hash(key)
	current := h.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}
