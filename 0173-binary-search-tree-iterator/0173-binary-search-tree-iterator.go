/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode // Tumpukan untuk melacak simpul-simpul yang akan diunjukkan
}

// Constructor digunakan untuk menginisialisasi objek BSTIterator
// Parameter "root" adalah akar dari pohon biner yang akan diiterasi
func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{} // Inisialisasi tumpukan kosong
	current := root

	// Inisialisasi tumpukan dengan simpul-simpul dalam urutan in-order
	for current != nil {
		stack = append(stack, current)
		current = current.Left
	}

	// Mengembalikan objek BSTIterator dengan tumpukan yang telah diisi
	return BSTIterator{stack}
}

// Next() digunakan untuk memindahkan iterator ke simpul selanjutnya
// dan mengembalikan nilai dari simpul tersebut
func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1] // Mengambil simpul teratas dari tumpukan
	this.stack = this.stack[:len(this.stack)-1] // Menghapus simpul teratas dari tumpukan

	// Jika simpul memiliki anak kanan, tambahkan anak kanan dan seluruh anak kiri dari anak kanan
	if node.Right != nil {
		current := node.Right
		for current != nil {
			this.stack = append(this.stack, current)
			current = current.Left
		}
	}

	return node.Val // Mengembalikan nilai dari simpul yang diambil
}

// HasNext() digunakan untuk memeriksa apakah masih ada simpul yang harus diunjukkan
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0 // Mengembalikan true jika tumpukan masih memiliki elemen
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */